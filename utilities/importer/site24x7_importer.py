import json
import os, sys
import subprocess
import logging

import argparse
 
WorkingDirectory = os.getcwd()

# ProjectHome = WorkingDirectory.replace("utilities", "")
# TerraformConfigurationFile = ProjectHome + "main.tf"
EmptyConfigurationFile = WorkingDirectory + os.path.sep + "empty_configuration.tf"
ImportedConfigurationFile = WorkingDirectory + os.path.sep + "output"+ os.path.sep +"imported_configuration.tf"
ImportCommandsFile = WorkingDirectory + os.path.sep + "output"+ os.path.sep +"import_commands.sh"
MonitorsToImportFile = WorkingDirectory + os.path.sep + "monitors_to_import.json"
Site24x7TerraformResourceNameVsAttributeTypesFile = WorkingDirectory + os.path.sep + "conf" + os.path.sep + "resource_vs_attribute_types.json"
TerraformStateFile = WorkingDirectory + os.path.sep + "terraform.tfstate"

MonitorsToImport = None
Site24x7TerraformResourceNameVsAttributeTypes = None
Site24x7TerraformResourceVsMonitorType = {
	"site24x7_website_monitor" : "URL",
	"site24x7_ssl_monitor" : "SSL_CERT",
	"site24x7_rest_api_monitor" : "RESTAPI",
	"site24x7_server_monitor" : "SERVER",
}

ResourceNameVsAttributesJSONInState = {}
ResourceNameVsFullConfiguration = {}


Site24x7TerraformResourceName = None
ResourceType = None

TfState = None
S247Importer = None
CommandLineParser = None


class Site24x7Importer:
	# Load monitors to import
	# Populate empty configuration needed for import
	def __init__(self, filePath):
		self.file = filePath
		self.monitors_to_import = None
		# Mandatory input from command line
		self.resource_type_in_site24x7 = Site24x7TerraformResourceVsMonitorType[Site24x7TerraformResourceName]
		self.resource_name_vs_empty_configuration = {}
		self.resource_name_vs_import_commands = {}
		self.site24x7_resource_vs_attribute_types = None
		self.load_monitors_to_import()
		self.load_site24x7_resource_vs_attribute_types()
		self.populate_empty_configuration()

	# Load monitors to import
	def load_monitors_to_import(self):
		if not os.path.exists(self.file):
			logging.info("Unable to find monitors to import file : "+self.file)
			sys.exit(1)
		try:
			self.monitors_to_import = FileUtil.read_json(self.file)
		except Exception as e:
			logging.info("Error while loading monitors_to_import.json : "+str(e))	
			logging.info("Please provide the list of monitors to be imported (eg) [\"123\", \"456\", \"789\"] in monitors_to_import.json ")
		logging.info("Monitors to import : "+str(self.monitors_to_import))

	def load_site24x7_resource_vs_attribute_types(self):
		if not os.path.exists(Site24x7TerraformResourceNameVsAttributeTypesFile):
			logging.info("Unable to find resource_vs_attribute_types.json file : ")
			sys.exit(1)
		try:
			self.site24x7_resource_vs_attribute_types = FileUtil.read_json(Site24x7TerraformResourceNameVsAttributeTypesFile)
		except Exception as e:
			logging.info("Error while loading resource_vs_attribute_types.json : "+str(e))	
		# logging.info("Monitors to import : "+str(self.site24x7_resource_vs_attribute_types))

	# Populate empty configuration needed for import
	def populate_empty_configuration(self):
		for monitorID in self.monitors_to_import:
			resourceName = ResourceType+"_"+monitorID
			resourceStr = "resource \""+Site24x7TerraformResourceName+"\" \""+resourceName+ "\" {"
			filebuffer = ["\n",resourceStr, "\n", "}","\n"]
			logging.debug("Empty Configuration : "+ ''.join(filebuffer))
			self.resource_name_vs_empty_configuration[resourceName] = ''.join(filebuffer)
			# Import Command
			import_command_list = ["terraform", "import", Site24x7TerraformResourceName+"."+resourceName, monitorID]
			self.resource_name_vs_import_commands[resourceName] = import_command_list

	def import_monitors(self):
		terraform_conf_file_data = FileUtil.read(EmptyConfigurationFile)
		for monitorID in self.monitors_to_import:
			resourceName = self.resource_type_in_site24x7+"_"+monitorID
			resource_name_vs_details_in_state = TfState.get_resource_name_vs_details()
			if resourceName not in resource_name_vs_details_in_state:
				logging.info("Importing the resource : "+resourceName)
				terraform_conf_file_data
				empty_conf_to_append = self.resource_name_vs_empty_configuration[resourceName]
				# Check whether the empty configuration is already present in the file.
				if empty_conf_to_append not in terraform_conf_file_data:
					FileUtil.append(EmptyConfigurationFile, empty_conf_to_append)
				# Execute the import command
				Util.execute_command(self.resource_name_vs_import_commands[resourceName])
			else:
				logging.info("Resource : "+resourceName+" already imported")
		# Parse the state file to populate resource_name_vs_attributes
		TfState.parse()

	def convert_attributes_from_state_file_to_configuration(self):
		resource_name_vs_attributes = TfState.get_resource_name_vs_attributes()
		print("resource_name_vs_attributes ============= ",resource_name_vs_attributes)
		if not resource_name_vs_attributes:
			logging.info("Failed to convert attributes from state file to configuration!! resource_name_vs_attributes info is empty in terraform state")
			return
		for resourceName in self.resource_name_vs_empty_configuration.keys():
			configurationBuffer = []
			resourceStr = "resource \""+Site24x7TerraformResourceName+"\" \""+resourceName+ "\" {"
			configurationBuffer.append("\n")
			configurationBuffer.append(resourceStr)
			attributesMap = resource_name_vs_attributes[resourceName]
			for attribute in attributesMap:
				# logging.info(attribute+" : "+str(attributesMap[attribute]))
				if attribute == "id":
					continue
				if attributesMap[attribute]:
					formatted_attribute = self.get_formatted_attribute(attribute, attributesMap)
					# logging.info("formatted_attribute : "+ str(formatted_attribute))
					if formatted_attribute:
						configurationBuffer.append(" \n ")
						configurationBuffer.append(attribute)
						configurationBuffer.append(" = ")
						configurationBuffer.append(formatted_attribute)
			configurationBuffer.append("\n}\n")
			ResourceNameVsFullConfiguration[resourceName] = self.get_formatted_configuration(configurationBuffer)
			logging.info("Configuration : "+ ResourceNameVsFullConfiguration[resourceName])

	def get_formatted_configuration(self, configurationBuffer):
		confStr = ''.join(map(str, configurationBuffer))
		confStr = confStr.replace("\'","\"")
		confStr = confStr.replace("True", "true")
		confStr = confStr.replace("False","false")
		# confStr = confStr.encode('ascii')
		return confStr

	def get_formatted_attribute(self, attribute, attributes_map):
		to_return = None
		attribute_type = self.get_attribute_type(attribute)
		# logging.info("attribute_type : "+ str(attribute_type))
		if attribute_type == "str":
			to_return = "\""+attributes_map[attribute]+"\""
		elif attribute_type == "list":
			to_return = [str(i) for i in attributes_map[attribute]]
		else:
			to_return = attributes_map[attribute]
		return to_return

	def get_attribute_type(self, attribute):
		attribute_name_vs_type = self.site24x7_resource_vs_attribute_types[Site24x7TerraformResourceName]
		if attribute in attribute_name_vs_type:
			return attribute_name_vs_type[attribute]

	def write_imported_configuration(self):
		config_str = ''
		for resourceName in ResourceNameVsFullConfiguration.keys():
			config = ResourceNameVsFullConfiguration[resourceName]
			config_str = config_str + config + "\n"
		FileUtil.write(ImportedConfigurationFile, config_str)
		logging.info("Please check the imported configuration in "+ImportedConfigurationFile)

	def write_import_commands(self):
		import_commands_list = []
		import_commands_list.append("#!/bin/bash")
		for import_command in self.resource_name_vs_import_commands.values():
			import_commands_list.append("\n")
			import_commands_list.append(" ".join(import_command))
		print("import_commands_list : ",import_commands_list)
		FileUtil.write(ImportCommandsFile, import_commands_list)
		logging.info("Please execute the import_commands.sh file for importing your monitors : "+ImportCommandsFile)	



class TerraformState:
	def __init__(self, filePath):
		logging.info("Loading Terraform state information")
		self.file = filePath
		self.resource_name_vs_details = {}
		self.resource_name_vs_attributes = {}
		self.parse()

	def get_resource_name_vs_details(self):
		return self.resource_name_vs_details

	def get_resource_name_vs_attributes(self):
		return self.resource_name_vs_attributes

	# Parses the terraform state file and populates resource_name_vs_attributes
	def parse(self):
		if not os.path.exists(self.file):
			logging.info("Unable to find the Terraform state file : "+self.file)
			return
		with open(TerraformStateFile, 'r') as terraformStateFile:
			stateFileJSON = json.load(terraformStateFile)
			if "resources" in stateFileJSON:
				resourcesList = stateFileJSON["resources"]
				for resource in resourcesList:
					if resource["mode"] == "managed":
						resourceName = resource["name"]
						self.resource_name_vs_details[resourceName] = resource
						self.populate_resource_name_vs_attributes_json(resourceName, resource)
		logging.info("Resources in Terraform State : "+str(self.resource_name_vs_details.keys()))

	def populate_resource_name_vs_attributes_json(self, resourceName, resourceMap):
		instancesList = resourceMap["instances"]
		for instanceMap in instancesList:
			if "attributes" in instanceMap:
				attributesMap = instanceMap["attributes"]
				self.resource_name_vs_attributes[resourceName] = attributesMap

	# Utility for grabbing attribute types
	def getAttribute_types(self):
		dict_to_return = {}
		for resourceName in self.resource_name_vs_attributes.keys():
			attributesDict = {}
			attributesMap = self.resource_name_vs_attributes[resourceName]
			for attribute in attributesMap:
				val = attributesMap[attribute]
				typeOfVal = type(val).__name__
				if typeOfVal == "NoneType":
					typeOfVal = "str"
				elif typeOfVal == "dict":
					typeOfVal = "map"
				attributesDict[attribute] = typeOfVal
			dict_to_return[resourceName] = attributesDict	
		return dict_to_return

	def write_attribute_types(self):
		attributeTypes = self.getAttribute_types()
		FileUtil.write_json(Site24x7TerraformResourceNameVsAttributeTypesFile, attributeTypes)


class FileUtil:
	@staticmethod
	def write_json(file_name, data):
		with open(file_name, "w") as file_handle:
			json_data = json.dumps(data)
			file_handle.write(json_data)
	@staticmethod
	def append(file_name, data_to_append):
		with open(file_name, 'a') as file_handle:
			file_handle.writelines(data_to_append)
	@staticmethod
	def write(file_name, data_to_append):
		with open(file_name, 'w') as file_handle:
			file_handle.writelines(data_to_append)
	@staticmethod
	def read_json(file_name):
		json_to_return = None
		with open(file_name, 'r') as file_handle:
			data=file_handle.read()
			json_to_return = json.loads(data)
		return json_to_return
	@staticmethod
	def read(file_name):
		data_to_return = None
		with open(file_name, 'r') as file_handle:
			data_to_return = file_handle.read()
		return data_to_return
		

class Util:
	@staticmethod
	def initialize_logging():
		logging.basicConfig(level=logging.INFO)
		logging.info("Logging Initialized : WorkingDirectory : " + WorkingDirectory)

	# terraform import site24x7_website_monitor.url_123456000026467038 123456000025786003
	@staticmethod
	def execute_command(command_list):
		process = subprocess.Popen(command_list, 
							stdout=subprocess.PIPE,
							# cwd=ProjectHome,
							universal_newlines=True)
		while True:
			output = process.stdout.readline()
			logging.info(output.strip())
			# Do something else
			return_code = process.poll()
			if return_code is not None:
				logging.info('RETURN CODE : '+ str(return_code))
				# Process has finished, read rest of the output 
				for output in process.stdout.readlines():
					logging.info(output.strip())
				break

	@staticmethod		
	def parse_command_line_args():
		global CommandLineParser, Site24x7TerraformResourceName, ResourceType
		CommandLineParser = argparse.ArgumentParser(description='Optional command description')
		CommandLineParser.add_argument('--resource', type=str, required=True, 
							help='Type of the Site24x7 terraform resource')

		args = CommandLineParser.parse_args()
		logging.info("Input argument resource : " + args.resource)
		Site24x7TerraformResourceName = args.resource
		ResourceType = Site24x7TerraformResourceVsMonitorType[Site24x7TerraformResourceName]		


def init():
	global TfState, S247Importer
	Util.initialize_logging()
	Util.parse_command_line_args()
	S247Importer = Site24x7Importer(MonitorsToImportFile)
	TfState = TerraformState(TerraformStateFile)


def main():
	init()
	# Invoke S247Importer.import_monitors() only after populating state information
	S247Importer.import_monitors()
	S247Importer.convert_attributes_from_state_file_to_configuration()
	S247Importer.write_imported_configuration()
	S247Importer.write_import_commands()


main()