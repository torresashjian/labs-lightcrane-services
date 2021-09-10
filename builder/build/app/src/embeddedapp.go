// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "app",
  "type": "flogo:app",
  "version": "1.0.0",
  "description": "",
  "appModel": "1.1.1",
  "imports": [
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filereader",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/trigger/execlistener",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectmaker",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/pipelinespliter",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/trigger/httpfilehandler",
    "github.com/project-flogo/flow",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filewriter",
    "github.com/project-flogo/flow/activity/subflow",
    "github.com/project-flogo/contrib/function/string",
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/project-flogo/legacybridge",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer",
    "github.com/project-flogo/contrib/function/array"
  ],
  "properties": [
    {
      "name": "System.ServiceLocatorIP",
      "type": "string",
      "value": "localhost"
    },
    {
      "name": "System.BaseFolder",
      "type": "string",
      "value": "/"
    },
    {
      "name": "System.FlogoBuilder",
      "type": "string",
      "value": "Flogo builder not set"
    }
  ],
  "triggers": [
    {
      "id": "ReceiveHTTPMessage",
      "ref": "#rest",
      "settings": {
        "certFile": "",
        "enableTLS": false,
        "keyFile": "",
        "port": 10083
      },
      "handlers": [
        {
          "name": "Build_Service",
          "settings": {
            "method": "POST",
            "path": "/f1/builder/build/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Build_Service"
              },
              "input": {
                "AppDescriptor": "=$.content.Descriptor",
                "ComponentType": "=$.content.ComponentType",
                "DeployDescriptor": {
                  "mapping": {
                    "Descriptor": "=$.content.DeployDescriptor",
                    "Name": "=$.pathParams.name",
                    "Properties": "=$.content.Properties",
                    "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                    "Type": "=$.content.ServiceType",
                    "Variables": "=$.content.Variables"
                  }
                },
                "Image": "=$.content.Image",
                "ProjectID": "=$.pathParams.id",
                "Runner": "=$.content.Runner",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "List_Runner",
          "settings": {
            "method": "GET",
            "path": "/f1/builder/listRunner"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:List_Runner"
              },
              "input": {
                "Command": "list"
              },
              "output": {
                "data": "=$.Runner"
              }
            }
          ]
        },
        {
          "name": "Get_Runner",
          "settings": {
            "method": "GET",
            "path": "/f1/builder/getRunner/:ID"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Runner"
              },
              "input": {
                "ID": "=$.pathParams.ID"
              },
              "output": {
                "data": "=$.DataOut"
              }
            }
          ]
        },
        {
          "name": "Docker_Container",
          "settings": {
            "method": "POST",
            "path": "/f1/builder/buildDocker/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Docker_Container"
              },
              "input": {
                "AppDescriptor": "=$.content.AppDescriptor",
                "ComponentType": "=$.content.ComponentType",
                "DeployDescriptor": "=$.content.DeployDescriptor",
                "ID": "=$.pathParams.id",
                "Image": "=$.content.Image",
                "Runner": "=$.content.Runner",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "AWS_Lambda",
          "settings": {
            "method": "POST",
            "path": "/f1/builder/buildLambda/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:AWS_Lambda"
              },
              "input": {
                "AppDescriptor": "=$.content.AppDescriptor",
                "ComponentType": "=$.content.ComponentType",
                "DeployDescriptor": "=$.content.DeployDescriptor",
                "ID": "=$.pathParams.id",
                "Image": "=$.content.Image",
                "Runner": "=$.content.Runner",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "Upload_Resource",
          "settings": {
            "method": "PUT",
            "path": "/f1/builder/upload/:type/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Upload_Resource"
              },
              "input": {
                "ID": "=$.pathParams.id",
                "Path": "=$.content.Path",
                "Resource": "=$.content.Resource",
                "Type": "=$.pathParams.type"
              },
              "output": {
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ]
        }
      ]
    },
    {
      "id": "Timer",
      "ref": "#timer",
      "settings": {},
      "handlers": [
        {
          "name": "Ping",
          "settings": {
            "repeatInterval": "60s",
            "startDelay": ""
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Ping"
              },
              "input": {
                "Now": ""
              }
            }
          ]
        }
      ]
    },
    {
      "id": "HTTPFileHandler",
      "ref": "#httpfilehandler",
      "settings": {
        "Port": "10084"
      },
      "handlers": [
        {
          "name": "Push_And_Build_Project",
          "settings": {
            "BaseFolder": "/home/f1/tmp/",
            "Path": "/f1/builder/pushAndBuild/{compressedProject}"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Push_And_Build_Project"
              },
              "input": {
                "FilePath": "=$.FilePath",
                "Filename": "=$.Filename"
              },
              "output": {
                "data": "=$.OutputData"
              }
            }
          ],
          "schemas": {
            "reply": {
              "data": {
                "fe_metadata": "{}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
              }
            }
          }
        }
      ]
    },
    {
      "id": "ExecEventListener",
      "ref": "#execlistener",
      "settings": {},
      "handlers": [
        {
          "name": "Build_Event",
          "settings": {
            "eventBoker": "DockerBuildEvent"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Build_Event"
              },
              "input": {
                "Event": "=$.Event"
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:Generate_F1_Descriptor",
      "data": {
        "name": "Generate F1 Descriptor",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogGenerateF1Descriptor",
            "to": "ConstructF1Dscrptr",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ConstructF1Dscrptr",
            "to": "SerializeF1Dscrptr",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SerializeF1Dscrptr",
            "to": "WriterF1Descriptr",
            "type": "default"
          },
          {
            "id": 4,
            "from": "WriterF1Descriptr",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogGenerateF1Descriptor",
            "name": "LogGenerateF1Descriptor",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%%% Gnerate F1 Descriptor %%%%%% ID : \", $flow.ID, \", DescriptorFilename : \", $flow.DescriptorFilename, \", DeployDescriptor : \", coerce.toString($flow.DeployDescriptor))"
              }
            }
          },
          {
            "id": "ConstructF1Dscrptr",
            "name": "ConstructF1Dscrptr",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n\t\"DeploymentGpID\" : \"wafermaps-pipeline\",\n\t\"DataFlow\" : [{\n\t\t\t\"Upstream\" : \"hello\",\n\t\t\t\"Downstream\" : \"rpy2-wafermaps\"\n\t\t}\n\t],\n    \"Components\":[\n        {\n\t\t\t\"Type\":\"Source\",\n\t\t\t\"Runtime\":\"flogo\",\n\t\t\t\"Name\":\"s-0\",\n\t\t\t\"Replicas\": 1,\n\t\t\t\"DockerImage\" : \"bigoyang/inference-endpoint:0.1.0\",\n\t\t\t\"Volumes\" : [\n\t\t\t\t{\n\t\t\t\t\t\"Name\" : \"metadata\",\n\t\t\t\t\t\"MountPoint\" : \"/usr/local/Data\"\n\t\t\t\t}\n\t\t\t],\n            \"Properties\":[\n            \t{\n                \t\"Name\":\"\",\n                \t\"Value\":\"\",\n                \t\"Type\":\"\"\n\t\t\t\t}\n\t\t\t]\n        }\n\t],\n    \"Services\":[\n        {\n\t\t\t\"Type\":\"aws-lambda\",\n\t\t\t\"Name\":\"flogo\",\n\t\t\t\"Descriptor\":\"serverless.yml\",\n            \"Variables\" : {},\n\t\t\t\t\"Properties\": [{\n\t\t\t\t\"Group\" : \"\",\n\t\t\t\t\"Value\" : [{\n\t\t\t\t\t\"Name\": \"\",\n\t\t\t\t\t\"Value\": \"\",\n\t\t\t\t\t\"Type\": \"\"\n\t\t\t\t}]\n\t\t\t}]\n\t\t}\n\t],\n    \"System\" : {\n    }\n}\n"
                },
                "defaultValue": ""
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "Services": "=array.create($flow.DeployDescriptor)",
                    "DeploymentGpID": "=$flow.ID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"DeploymentGpID\":\"wafermaps-pipeline\",\"DataFlow\":[{\"Upstream\":\"hello\",\"Downstream\":\"rpy2-wafermaps\"}],\"Components\":[{\"Type\":\"Source\",\"Runtime\":\"flogo\",\"Name\":\"s-0\",\"Replicas\":1,\"DockerImage\":\"bigoyang/inference-endpoint:0.1.0\",\"Volumes\":[{\"Name\":\"metadata\",\"MountPoint\":\"/usr/local/Data\"}],\"Properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}],\"Services\":[{\"Type\":\"aws-lambda\",\"Name\":\"flogo\",\"Descriptor\":\"serverless.yml\",\"Variables\":{},\"Properties\":[{\"Group\":\"\",\"Value\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}]}],\"System\":{}}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"DeploymentGpID\":\"wafermaps-pipeline\",\"DataFlow\":[{\"Upstream\":\"hello\",\"Downstream\":\"rpy2-wafermaps\"}],\"Components\":[{\"Type\":\"Source\",\"Runtime\":\"flogo\",\"Name\":\"s-0\",\"Replicas\":1,\"DockerImage\":\"bigoyang/inference-endpoint:0.1.0\",\"Volumes\":[{\"Name\":\"metadata\",\"MountPoint\":\"/usr/local/Data\"}],\"Properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}],\"Services\":[{\"Type\":\"aws-lambda\",\"Name\":\"flogo\",\"Descriptor\":\"serverless.yml\",\"Variables\":{},\"Properties\":[{\"Group\":\"\",\"Value\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}]}],\"System\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n\t\"DeploymentGpID\" : \"wafermaps-pipeline\",\n\t\"DataFlow\" : [{\n\t\t\t\"Upstream\" : \"hello\",\n\t\t\t\"Downstream\" : \"rpy2-wafermaps\"\n\t\t}\n\t],\n    \"Components\":[\n        {\n\t\t\t\"Type\":\"Source\",\n\t\t\t\"Runtime\":\"flogo\",\n\t\t\t\"Name\":\"s-0\",\n\t\t\t\"Replicas\": 1,\n\t\t\t\"DockerImage\" : \"bigoyang/inference-endpoint:0.1.0\",\n\t\t\t\"Volumes\" : [\n\t\t\t\t{\n\t\t\t\t\t\"Name\" : \"metadata\",\n\t\t\t\t\t\"MountPoint\" : \"/usr/local/Data\"\n\t\t\t\t}\n\t\t\t],\n            \"Properties\":[\n            \t{\n                \t\"Name\":\"\",\n                \t\"Value\":\"\",\n                \t\"Type\":\"\"\n\t\t\t\t}\n\t\t\t]\n        }\n\t],\n    \"Services\":[\n        {\n\t\t\t\"Type\":\"aws-lambda\",\n\t\t\t\"Name\":\"flogo\",\n\t\t\t\"Descriptor\":\"serverless.yml\",\n            \"Variables\" : {},\n\t\t\t\t\"Properties\": [{\n\t\t\t\t\"Group\" : \"\",\n\t\t\t\t\"Value\" : [{\n\t\t\t\t\t\"Name\": \"\",\n\t\t\t\t\t\"Value\": \"\",\n\t\t\t\t\t\"Type\": \"\"\n\t\t\t\t}]\n\t\t\t}]\n\t\t}\n\t],\n    \"System\" : {\n    }\n}\n"
                  }
                }
              }
            }
          },
          {
            "id": "SerializeF1Dscrptr",
            "name": "SerializeF1Dscrptr",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[ConstructF1Dscrptr].ObjectOut"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "WriterF1Descriptr",
            "name": "WriterF1Descriptr",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$Filename$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[SerializeF1Dscrptr].SerializedString"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$flow.DescriptorFilename"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "ErrorCode": 100,
                  "Message": "Descriptor Generated!"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::Generate F1 Descriptor - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "DescriptorFilename",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:Build_General_Pipeline",
      "data": {
        "name": "Build General Pipeline",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "PipelineSpliter",
            "to": "ProjectProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ProjectProperties",
            "to": "JSONSerializer",
            "type": "default"
          },
          {
            "id": 3,
            "from": "JSONSerializer",
            "to": "SaveDeployDescriptor",
            "type": "default"
          },
          {
            "id": 4,
            "from": "SaveDeployDescriptor",
            "to": "Return",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ProjectProperties",
            "to": "BuildFlogoAppFilePath",
            "type": "label",
            "label": "Mapping to SaveFlogoComponents"
          },
          {
            "id": 6,
            "from": "BuildFlogoAppFilePath",
            "to": "SaveFlogoComponents",
            "type": "default"
          },
          {
            "id": 7,
            "from": "SaveFlogoComponents",
            "to": "Log",
            "type": "default"
          },
          {
            "id": 8,
            "from": "Log",
            "to": "BuildFlogoApp",
            "type": "default"
          },
          {
            "id": 9,
            "from": "ProjectProperties",
            "to": "BuildExtApp",
            "type": "expression",
            "label": "MappingtoBuildExtApp",
            "value": "false"
          }
        ],
        "tasks": [
          {
            "id": "PipelineSpliter",
            "name": "PipelineSpliter",
            "description": "This activity split flogo flow pipeline",
            "activity": {
              "ref": "#pipelinespliter",
              "settings": {
                "Template": {
                  "filename": "PipelineTemplate.json",
                  "content": "data:application/json;base64,ewogICAgImltcG9ydHMiOiBbCiAgICAgICAgImdpdGh1Yi5jb20vcHJvamVjdC1mbG9nby9mbG93IiwKICAgICAgICAiZ2l0aHViLmNvbS9USUJDT1NvZnR3YXJlL0dyYXBoQnVpbGRlcl9Ub29scy9hY3Rpdml0eS9qc29uZGVzZXJpYWxpemVyIiwKICAgICAgICAiZ2l0aHViLmNvbS9USUJDT1NvZnR3YXJlL0dyYXBoQnVpbGRlcl9Ub29scy9hY3Rpdml0eS9qc29uc2VyaWFsaXplciIsCiAgICAgICAgImdpdGh1Yi5jb20vVElCQ09Tb2Z0d2FyZS9Nb2RlbE9wcy9hY3Rpdml0eS9waXBlY291cGxlciIsCiAgICAgICAgImdpdGh1Yi5jb20vcHJvamVjdC1mbG9nby9jb250cmliL2FjdGl2aXR5L2FjdHJldHVybiIsCiAgICAgICAgImdpdGh1Yi5jb20vcHJvamVjdC1mbG9nby9ncnBjL3RyaWdnZXIvZ3JwYyIKICAgIF0sCiAgICAibmFtZSI6ICJQaXBlbGluZVRlbXBsYXRlIiwKICAgICJkZXNjcmlwdGlvbiI6ICIgIiwKICAgICJ2ZXJzaW9uIjogIjEuMS4wIiwKICAgICJ0eXBlIjogImZsb2dvOmFwcCIsCiAgICAiYXBwTW9kZWwiOiAiMS4xLjEiLAogICAgImZlVmVyc2lvbiI6ICIyLjguMSIsCiAgICAidHJpZ2dlcnMiOiBbCiAgICAgICAgewogICAgICAgICAgICAicmVmIjogIiNncnBjIiwKICAgICAgICAgICAgIm5hbWUiOiAiZ3JwYy10cmlnZ2VyIiwKICAgICAgICAgICAgImRlc2NyaXB0aW9uIjogImdSUEMgVHJpZ2dlciIsCiAgICAgICAgICAgICJzZXR0aW5ncyI6IHsKICAgICAgICAgICAgICAgICJwb3J0IjogIj0kcHJvcGVydHlbXCJwaXBlY291cGxlci5wb3J0XCJdIiwKICAgICAgICAgICAgICAgICJwcm90b05hbWUiOiAicGlwZWNvdXBsZXIucHJvdG8iLAogICAgICAgICAgICAgICAgInByb3RvRmlsZSI6IHsKICAgICAgICAgICAgICAgICAgICAiZmlsZW5hbWUiOiAicGlwZWNvdXBsZXIucHJvdG8iLAogICAgICAgICAgICAgICAgICAgICJjb250ZW50IjogImRhdGE6YXBwbGljYXRpb24vb2N0ZXQtc3RyZWFtO2Jhc2U2NCxjM2x1ZEdGNElEMGdJbkJ5YjNSdk15STdDZ3B2Y0hScGIyNGdhbUYyWVY5dGRXeDBhWEJzWlY5bWFXeGxjeUE5SUhSeWRXVTdDbTl3ZEdsdmJpQnFZWFpoWDNCaFkydGhaMlVnUFNBaVkyOXRMblJwWW1OdkxtMXZaR1ZzYjNCeklqc0tiM0IwYVc5dUlHcGhkbUZmYjNWMFpYSmZZMnhoYzNOdVlXMWxJRDBnSWxCcGNHVkRiM1Z3YkdWeVVISnZkRzhpT3dvS2NHRmphMkZuWlNCd2FYQmxZMjkxY0d4bGNqc0tDbk5sY25acFkyVWdVR2x3WlVOdmRYQnNaWElnZXdvSmNuQmpJRWhoYm1Sc1pVUmhkR0VnS0VSaGRHRXBJSEpsZEhWeWJuTWdLRkpsY0d4NUtTQjdmUXA5Q2dwdFpYTnpZV2RsSUVSaGRHRWdld29KYzNSeWFXNW5JSE5sYm1SbGNpQTlJREU3Q2dsemRISnBibWNnU1VRZ1BTQXlPd29KYzNSeWFXNW5JR052Ym5SbGJuUWdQU0F6T3dwOUNncHRaWE56WVdkbElGSmxjR3g1SUhzS0NYTjBjbWx1WnlCelpXNWtaWElnUFNBeE93b0pjM1J5YVc1bklFbEVJRDBnTWpzS0NYTjBjbWx1WnlCamIyNTBaVzUwSUQwZ016c0tDV0p2YjJ3Z2MzUmhkSFZ6SUQwZ05Ec0tmUW89IgogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICJlbmFibGVUTFMiOiBmYWxzZSwKICAgICAgICAgICAgICAgICJzZXJ2ZXJDZXJ0IjogIiIsCiAgICAgICAgICAgICAgICAic2VydmVyS2V5IjogIiIKICAgICAgICAgICAgfSwKICAgICAgICAgICAgImlkIjogImdSUENUcmlnZ2VyIiwKICAgICAgICAgICAgImhhbmRsZXJzIjogWwogICAgICAgICAgICAgICAgewogICAgICAgICAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICIiLAogICAgICAgICAgICAgICAgICAgICJzZXR0aW5ncyI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInNlcnZpY2VOYW1lIjogIlBpcGVDb3VwbGVyIiwKICAgICAgICAgICAgICAgICAgICAgICAgIm1ldGhvZE5hbWUiOiAiSGFuZGxlRGF0YSIKICAgICAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgICAgICJhY3Rpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJyZWYiOiAiZ2l0aHViLmNvbS9wcm9qZWN0LWZsb2dvL2Zsb3ciLAogICAgICAgICAgICAgICAgICAgICAgICAic2V0dGluZ3MiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAiZmxvd1VSSSI6ICJyZXM6Ly9mbG93OlBpcGVsaW5lX0Zsb3ciCiAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICJpbnB1dCI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJEYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJtYXBwaW5nIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiQ29udGVudCI6ICI9JC5wYXJhbXMuY29udGVudCIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJJRCI6ICI9JC5wYXJhbXMuSUQiLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiU2VuZGVyIjogIj0kLnBhcmFtcy5zZW5kZXIiCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgICAgICAgICAib3V0cHV0IjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgImRhdGEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIm1hcHBpbmciOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzZW5kZXIiOiAiPSQuUmVwbHkuU2VuZGVyIiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIklEIjogIj0kLlJlcGx5LklEIiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgImNvbnRlbnQiOiAiPSQuUmVwbHkuQ29udGVudCIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzdGF0dXMiOiAiPSQuUmVwbHkuU3RhdHVzIgogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgInJlcGx5IjogewogICAgICAgICAgICAgICAgICAgICAgICAiY29kZSI6IDIwMAogICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgInNjaGVtYXMiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJyZXBseSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJkYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImpzb24iLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ2YWx1ZSI6ICJ7XCJ0eXBlXCI6XCJvYmplY3RcIixcInByb3BlcnRpZXNcIjp7XCJzZW5kZXJcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJJRFwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcImNvbnRlbnRcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJzdGF0dXNcIjp7XCJ0eXBlXCI6XCJib29sZWFuXCJ9fSxcInJlcXVpcmVkXCI6W119IiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiZmVfbWV0YWRhdGEiOiAie1widHlwZVwiOlwib2JqZWN0XCIsXCJwcm9wZXJ0aWVzXCI6e1wic2VuZGVyXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiSURcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJjb250ZW50XCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwic3RhdHVzXCI6e1widHlwZVwiOlwiYm9vbGVhblwifX0sXCJyZXF1aXJlZFwiOltdfSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICAgICAgICAgIm91dHB1dCI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJwYXJhbXMiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInR5cGUiOiAianNvbiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInZhbHVlIjogIntcInR5cGVcIjpcIm9iamVjdFwiLFwicHJvcGVydGllc1wiOntcInNlbmRlclwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIklEXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiY29udGVudFwiOntcInR5cGVcIjpcInN0cmluZ1wifX0sXCJyZXF1aXJlZFwiOltdfSIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgImZlX21ldGFkYXRhIjogIntcInR5cGVcIjpcIm9iamVjdFwiLFwicHJvcGVydGllc1wiOntcInNlbmRlclwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIklEXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiY29udGVudFwiOntcInR5cGVcIjpcInN0cmluZ1wifX0sXCJyZXF1aXJlZFwiOltdfSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgXQogICAgICAgIH0KICAgIF0sCiAgICAicmVzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICAgImlkIjogImZsb3c6UGlwZWxpbmVfRmxvdyIsCiAgICAgICAgICAgICJkYXRhIjogewogICAgICAgICAgICAgICAgIm5hbWUiOiAiUGlwZWxpbmUgRmxvdyIsCiAgICAgICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiIiwKICAgICAgICAgICAgICAgICJsaW5rcyI6IFsKICAgICAgICAgICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAgICAgICAgICJpZCI6IDEsCiAgICAgICAgICAgICAgICAgICAgICAgICJmcm9tIjogIkpTT05EZXNlcmlhbGl6ZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAidG8iOiAiSlNPTlNlcmlhbGl6ZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAidHlwZSI6ICJkZWZhdWx0IgogICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgewogICAgICAgICAgICAgICAgICAgICAgICAiaWQiOiAyLAogICAgICAgICAgICAgICAgICAgICAgICAiZnJvbSI6ICJKU09OU2VyaWFsaXplciIsCiAgICAgICAgICAgICAgICAgICAgICAgICJ0byI6ICJNTHBpcGVsaW5lY291cGxlciIsCiAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImRlZmF1bHQiCiAgICAgICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAgICAgICAgICJpZCI6IDMsCiAgICAgICAgICAgICAgICAgICAgICAgICJmcm9tIjogIk1McGlwZWxpbmVjb3VwbGVyIiwKICAgICAgICAgICAgICAgICAgICAgICAgInRvIjogIlJldHVybiIsCiAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImRlZmF1bHQiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgXSwKICAgICAgICAgICAgICAgICJ0YXNrcyI6IFsKICAgICAgICAgICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAgICAgICAgICJpZCI6ICJKU09ORGVzZXJpYWxpemVyIiwKICAgICAgICAgICAgICAgICAgICAgICAgIm5hbWUiOiAiSlNPTkRlc2VyaWFsaXplciIsCiAgICAgICAgICAgICAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJKU09OIERlc2VyaWFsaXplciBBY3Rpdml0eSIsCiAgICAgICAgICAgICAgICAgICAgICAgICJhY3Rpdml0eSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJyZWYiOiAiI2pzb25kZXNlcmlhbGl6ZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgInNldHRpbmdzIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzYW1wbGUiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJmaWxlbmFtZSI6ICJpcmlzLmpzb24iLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiY29udGVudCI6ICJkYXRhOmFwcGxpY2F0aW9uL2pzb247YmFzZTY0LGV5QWljMlZ3WVd4ZmJHVnVaM1JvSWpvZ05TNHlMQ0FpYzJWd1lXeGZkMmxrZEdnaU9pQXpMallzSUNKd1pYUmhiRjlzWlc1bmRHZ2lPaUF4TGpRc0lDSndaWFJoYkY5M2FXUjBhQ0k2SURBdU15QjkiCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiZGVmYXVsdFZhbHVlIjogIiIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAiaW5wdXQiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIkpTT05TdHJpbmciOiAiPSRmbG93LkRhdGEuQ29udGVudCIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2NoZW1hcyI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAib3V0cHV0IjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiRGF0YSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImpzb24iLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInZhbHVlIjogIntcIiRzY2hlbWFcIjpcImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDQvc2NoZW1hI1wiLFwidHlwZVwiOlwib2JqZWN0XCIsXCJwcm9wZXJ0aWVzXCI6e1wic2VwYWxfbGVuZ3RoXCI6e1widHlwZVwiOlwibnVtYmVyXCJ9LFwic2VwYWxfd2lkdGhcIjp7XCJ0eXBlXCI6XCJudW1iZXJcIn0sXCJwZXRhbF9sZW5ndGhcIjp7XCJ0eXBlXCI6XCJudW1iZXJcIn0sXCJwZXRhbF93aWR0aFwiOntcInR5cGVcIjpcIm51bWJlclwifX19IiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJmZV9tZXRhZGF0YSI6ICJ7XCJzZXBhbF9sZW5ndGhcIjo1LjIsXCJzZXBhbF93aWR0aFwiOjMuNixcInBldGFsX2xlbmd0aFwiOjEuNCxcInBldGFsX3dpZHRoXCI6MC4zfSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgewogICAgICAgICAgICAgICAgICAgICAgICAiaWQiOiAiSlNPTlNlcmlhbGl6ZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAibmFtZSI6ICJKU09OU2VyaWFsaXplciIsCiAgICAgICAgICAgICAgICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJKU09OIFNlcmlhbGl6ZXIgQWN0aXZpdHkiLAogICAgICAgICAgICAgICAgICAgICAgICAiYWN0aXZpdHkiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAicmVmIjogIiNqc29uc2VyaWFsaXplciIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2V0dGluZ3MiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInNhbXBsZSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgImZpbGVuYW1lIjogImlyaXMuanNvbiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJjb250ZW50IjogImRhdGE6YXBwbGljYXRpb24vanNvbjtiYXNlNjQsZXlBaWMyVndZV3hmYkdWdVozUm9Jam9nTlM0eUxDQWljMlZ3WVd4ZmQybGtkR2dpT2lBekxqWXNJQ0p3WlhSaGJGOXNaVzVuZEdnaU9pQXhMalFzSUNKd1pYUmhiRjkzYVdSMGFDSTZJREF1TXlCOSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgICAgICAgICAgICAgImlucHV0IjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJEYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAibWFwcGluZyI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzZXBhbF9sZW5ndGgiOiAiPSRhY3Rpdml0eVtKU09ORGVzZXJpYWxpemVyXS5EYXRhLnNlcGFsX2xlbmd0aCIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2VwYWxfd2lkdGgiOiAiPSRhY3Rpdml0eVtKU09ORGVzZXJpYWxpemVyXS5EYXRhLnNlcGFsX3dpZHRoIiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJwZXRhbF9sZW5ndGgiOiAiPSRhY3Rpdml0eVtKU09ORGVzZXJpYWxpemVyXS5EYXRhLnBldGFsX2xlbmd0aCIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAicGV0YWxfd2lkdGgiOiAiPSRhY3Rpdml0eVtKU09ORGVzZXJpYWxpemVyXS5EYXRhLnBldGFsX3dpZHRoIgogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzY2hlbWFzIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJpbnB1dCI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIkRhdGEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAidHlwZSI6ICJqc29uIiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ2YWx1ZSI6ICJ7XCIkc2NoZW1hXCI6XCJodHRwOi8vanNvbi1zY2hlbWEub3JnL2RyYWZ0LTA0L3NjaGVtYSNcIixcInR5cGVcIjpcIm9iamVjdFwiLFwicHJvcGVydGllc1wiOntcInNlcGFsX2xlbmd0aFwiOntcInR5cGVcIjpcIm51bWJlclwifSxcInNlcGFsX3dpZHRoXCI6e1widHlwZVwiOlwibnVtYmVyXCJ9LFwicGV0YWxfbGVuZ3RoXCI6e1widHlwZVwiOlwibnVtYmVyXCJ9LFwicGV0YWxfd2lkdGhcIjp7XCJ0eXBlXCI6XCJudW1iZXJcIn19fSIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiZmVfbWV0YWRhdGEiOiAie1wic2VwYWxfbGVuZ3RoXCI6NS4yLFwic2VwYWxfd2lkdGhcIjozLjYsXCJwZXRhbF9sZW5ndGhcIjoxLjQsXCJwZXRhbF93aWR0aFwiOjAuM30iCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICAgICAgICAgImlkIjogIk1McGlwZWxpbmVjb3VwbGVyIiwKICAgICAgICAgICAgICAgICAgICAgICAgIm5hbWUiOiAiTUxwaXBlbGluZWNvdXBsZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhpcyBhY3Rpdml0eSBjb3VwbGUgTUwgcGlwZWxpbmVzIiwKICAgICAgICAgICAgICAgICAgICAgICAgImFjdGl2aXR5IjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgInJlZiI6ICIjcGlwZWNvdXBsZXIiLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgInNldHRpbmdzIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJEb3duc3RyZWFtSG9zdCI6ICI9JHByb3BlcnR5W1wicGlwZWNvdXBsZXIuZG93bnN0cmVhbUhvc3RzXCJdIiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiUG9ydCI6ICI9JHByb3BlcnR5W1wicGlwZWNvdXBsZXIucG9ydFwiXSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAiaW5wdXQiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIkRhdGEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJtYXBwaW5nIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIkNvbnRlbnQiOiAiPSRhY3Rpdml0eVtKU09OU2VyaWFsaXplcl0uSlNPTlN0cmluZyIKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2NoZW1hcyI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiaW5wdXQiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJEYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInR5cGUiOiAianNvbiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAidmFsdWUiOiAie1wiJHNjaGVtYVwiOlwiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNC9zY2hlbWEjXCIsXCJ0eXBlXCI6XCJvYmplY3RcIixcInByb3BlcnRpZXNcIjp7XCJTZW5kZXJcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJJRFwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIkNvbnRlbnRcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn19fSIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiZmVfbWV0YWRhdGEiOiAie1wiU2VuZGVyXCI6XCIyXCIsXCJJRFwiOlwiMlwiLFwiQ29udGVudFwiOlwiMlwifSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIm91dHB1dCI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIlJlcGx5IjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInR5cGUiOiAianNvbiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAidmFsdWUiOiAie1wiJHNjaGVtYVwiOlwiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNC9zY2hlbWEjXCIsXCJ0eXBlXCI6XCJvYmplY3RcIixcInByb3BlcnRpZXNcIjp7XCJTZW5kZXJcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJJZFwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIkNvbnRlbnRcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJTdGF0dXNcIjp7XCJ0eXBlXCI6XCJib29sZWFuXCJ9fX0iLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgImZlX21ldGFkYXRhIjogIntcIlNlbmRlclwiOlwiMlwiLFwiSWRcIjpcIjJcIixcIkNvbnRlbnRcIjpcIjJcIixcIlN0YXR1c1wiOnRydWV9IgogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAgICAgICAgICJpZCI6ICJSZXR1cm4iLAogICAgICAgICAgICAgICAgICAgICAgICAibmFtZSI6ICJSZXR1cm4iLAogICAgICAgICAgICAgICAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiU2ltcGxlIFJldHVybiBBY3Rpdml0eSIsCiAgICAgICAgICAgICAgICAgICAgICAgICJhY3Rpdml0eSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJyZWYiOiAiI2FjdHJldHVybiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2V0dGluZ3MiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIm1hcHBpbmdzIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAiUmVwbHkiOiAiPSRhY3Rpdml0eVtNTHBpcGVsaW5lY291cGxlcl0uUmVwbHkiCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgXSwKICAgICAgICAgICAgICAgICJtZXRhZGF0YSI6IHsKICAgICAgICAgICAgICAgICAgICAiaW5wdXQiOiBbCiAgICAgICAgICAgICAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJuYW1lIjogIkRhdGEiLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJzY2hlbWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInR5cGUiOiAianNvbiIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgInZhbHVlIjogIntcIlNlbmRlclwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIklEXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiQ29udGVudFwiOntcInR5cGVcIjpcInN0cmluZ1wifX0iCiAgICAgICAgICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICBdLAogICAgICAgICAgICAgICAgICAgICJvdXRwdXQiOiBbCiAgICAgICAgICAgICAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJuYW1lIjogIlJlcGx5IiwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogIm9iamVjdCIsCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAic2NoZW1hIjogewogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImpzb24iLAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ2YWx1ZSI6ICJ7XCJTZW5kZXJcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJJRFwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIkNvbnRlbnRcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJTdGF0dXNcIjp7XCJ0eXBlXCI6XCJib29sZWFuXCJ9fSIKICAgICAgICAgICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgIF0sCiAgICAgICAgICAgICAgICAgICAgImZlX21ldGFkYXRhIjogewogICAgICAgICAgICAgICAgICAgICAgICAiaW5wdXQiOiAie1wiJHNjaGVtYVwiOlwiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNC9zY2hlbWEjXCIsXCJ0eXBlXCI6XCJvYmplY3RcIixcInByb3BlcnRpZXNcIjp7XCJEYXRhXCI6e1widHlwZVwiOlwib2JqZWN0XCIsXCJwcm9wZXJ0aWVzXCI6e1wiU2VuZGVyXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiSURcIjp7XCJ0eXBlXCI6XCJzdHJpbmdcIn0sXCJDb250ZW50XCI6e1widHlwZVwiOlwic3RyaW5nXCJ9fX19fSIsCiAgICAgICAgICAgICAgICAgICAgICAgICJvdXRwdXQiOiAie1wiJHNjaGVtYVwiOlwiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNC9zY2hlbWEjXCIsXCJ0eXBlXCI6XCJvYmplY3RcIixcInByb3BlcnRpZXNcIjp7XCJSZXBseVwiOntcInR5cGVcIjpcIm9iamVjdFwiLFwicHJvcGVydGllc1wiOntcIlNlbmRlclwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIklEXCI6e1widHlwZVwiOlwic3RyaW5nXCJ9LFwiQ29udGVudFwiOntcInR5cGVcIjpcInN0cmluZ1wifSxcIlN0YXR1c1wiOntcInR5cGVcIjpcImJvb2xlYW5cIn19fX19IgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgIH0KICAgIF0sCiAgICAicHJvcGVydGllcyI6IFsKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogInBpcGVsaW5lLmRlc2NyaXB0b3IiLAogICAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAidmFsdWUiOiAiXCJwbGVhc2Ugb3ZlcnJpZGUgaXRcIiIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGlwZWNvdXBsZXIucG9ydCIsCiAgICAgICAgICAgICJ0eXBlIjogImZsb2F0NjQiLAogICAgICAgICAgICAidmFsdWUiOiA5OTk3CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogInBpcGVjb3VwbGVyLmRvd25zdHJlYW1Ib3N0cyIsCiAgICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJcInBsZWFzZSBvdmVycmlkZSBpdFwiIgogICAgICAgIH0KICAgIF0sCiAgICAiY29ubmVjdGlvbnMiOiB7fSwKICAgICJjb250cmliIjogIlczc2ljbVZtSWpvaVoybDBhSFZpTG1OdmJTOXdjbTlxWldOMExXWnNiMmR2TDJkeWNHTXZkSEpwWjJkbGNpOW5jbkJqSWl3aWN6TnNiMk5oZEdsdmJpSTZJbnRWVTBWU1NVUjlMMFJsWm1GMWJIUXZkSEpwWjJkbGNpOW5jbkJqSW4wc2V5SnlaV1lpT2lKbmFYUm9kV0l1WTI5dEwxUkpRa05QVTI5bWRIZGhjbVV2VFc5a1pXeFBjSE1pTENKek0yeHZZMkYwYVc5dUlqb2llMVZUUlZKSlJIMHZUVzlrWld4UGNITWlmU3g3SW5KbFppSTZJbWRwZEdoMVlpNWpiMjB2VkVsQ1EwOVRiMlowZDJGeVpTOUhjbUZ3YUVKMWFXeGtaWEpmVkc5dmJITWlMQ0p6TTJ4dlkyRjBhVzl1SWpvaWUxVlRSVkpKUkgwdlIzSmhjR2hDZFdsc1pHVnlYMVJ2YjJ4ekluMWQiLAogICAgImZlX21ldGFkYXRhIjogIlVFc0RCQW9BQUFBSUFMS1M5MUNycVNSR0ZRQUFBQk1BQUFBSUFBQUFZWEJ3TG1wemIyNnJWaW9weWt4UFR5MXl5OGt2TDFheWlvNnRCUUJRU3dFQ0ZBQUtBQUFBQ0FDeWt2ZFFxNmtrUmhVQUFBQVRBQUFBQ0FBQUFBQUFBQUFBQUFBQUFBQUFBQUFBWVhCd0xtcHpiMjVRU3dVR0FBQUFBQUVBQVFBMkFBQUFPd0FBQUFBQSIKfQ=="
                }
              },
              "input": {
                "RawPipelineConfig": "=$flow.AppDescriptor"
              },
              "schemas": {
                "output": {
                  "DataFlow": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Upstream\":\"2\",\"Downstream\":\"2\"}]"
                  },
                  "PipelineConfig": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}",
                    "fe_metadata": "[{\"Name\":\"2\",\"Type\":\"2\",\"ComponentConfig\":\"2\",\"Properties\":[{\"Name\":\"2\",\"Value\":\"2\"}]}]"
                  }
                }
              }
            }
          },
          {
            "id": "ProjectProperties",
            "name": "ProjectProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectFolder": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "JSONSerializer",
            "name": "JSONSerializer",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": {
                  "filename": "Pipeline.json",
                  "content": "data:application/json;base64,ewoJIklEIjogIiIsCgkiRGF0YUZsb3ciOiBbewoJCQkiVXBzdHJlYW0iOiAiIiwKCQkJIkRvd25zdHJlYW0iOiAiIgoJCX0KCV0sCgkiUGlwZWxpbmVDb25maWciOiBbewoJCQkiTmFtZSI6ICIiLAoJCQkiVHlwZSI6ICIiLAoJCQkiQ29tcG9uZW50Q29uZmlnIjogIiIsCgkJCSJQcm9wZXJ0aWVzIjogW3sKCQkJCQkiTmFtZSI6ICIiLAoJCQkJCSJWYWx1ZSI6ICIiCgkJCQl9CgkJCV0KCQl9CgldCn0K"
                },
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"PipelineConfig\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n\t\"ID\": \"\",\n\t\"DataFlow\": [{\n\t\t\t\"Upstream\": \"\",\n\t\t\t\"Downstream\": \"\"\n\t\t}\n\t],\n\t\"PipelineConfig\": [{\n\t\t\t\"Name\": \"\",\n\t\t\t\"Type\": \"\",\n\t\t\t\"ComponentConfig\": \"\",\n\t\t\t\"Properties\": [{\n\t\t\t\t\t\"Name\": \"\",\n\t\t\t\t\t\"Value\": \"\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "ID": "=$activity[PipelineSpliter].ID",
                    "DataFlow": {
                      "@foreach($activity[PipelineSpliter].DataFlow,DataFlow)": {
                        "=": "$loop"
                      }
                    },
                    "PipelineConfig": {
                      "@foreach($activity[PipelineSpliter].PipelineConfig,PipelineConfig)": {
                        "=": "$loop"
                      }
                    }
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"PipelineConfig\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}}",
                    "fe_metadata": "{\"ID\":\"\",\"DataFlow\":[{\"Upstream\":\"\",\"Downstream\":\"\"}],\"PipelineConfig\":[{\"Name\":\"\",\"Type\":\"\",\"ComponentConfig\":\"\",\"Properties\":[{\"Name\":\"\",\"Value\":\"\"}]}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"PipelineConfig\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}}",
                    "fe_metadata": "{\n\t\"ID\": \"\",\n\t\"DataFlow\": [{\n\t\t\t\"Upstream\": \"\",\n\t\t\t\"Downstream\": \"\"\n\t\t}\n\t],\n\t\"PipelineConfig\": [{\n\t\t\t\"Name\": \"\",\n\t\t\t\"Type\": \"\",\n\t\t\t\"ComponentConfig\": \"\",\n\t\t\t\"Properties\": [{\n\t\t\t\t\t\"Name\": \"\",\n\t\t\t\t\t\"Value\": \"\"\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}\n"
                  }
                }
              }
            }
          },
          {
            "id": "SaveDeployDescriptor",
            "name": "SaveDeployDescriptor",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "%ProjectFolder%/descriptor/deploy.descriptor",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[JSONSerializer].JSONString"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectFolder": "=$activity[ProjectProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "ErrorCode": 200,
                  "Message": "=$activity[JSONSerializer].JSONString"
                }
              }
            }
          },
          {
            "id": "BuildFlogoAppFilePath",
            "name": "BuildFlogoAppFilePath",
            "description": "Mapping field from input to output",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[PipelineSpliter].PipelineConfig",
              "accumulate": true
            },
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"FlogoAppFilePath\",\"Type\":\"String\"},{\"Name\":\"ComponentConfig\",\"Type\":\"String\"},{\"Name\":\"AppType\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "FlogoAppFilePath": "=$activity[ProjectProperties].Data.ProjectFolder+\"/flogo/\"+$iteration[value].Name+\".json\"",
                    "ComponentConfig": "=$iteration[value].ComponentConfig",
                    "AppType": "=$iteration[value].Type"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoAppFilePath\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"AppType\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"FlogoAppFilePath\":\"2\",\"ComponentConfig\":\"2\",\"AppType\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoAppFilePath\":{\"type\":\"string\"},\"ComponentConfig\":{\"type\":\"string\"},\"AppType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"FlogoAppFilePath\":\"2\",\"ComponentConfig\":\"2\",\"AppType\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"FlogoAppFilePath\",\"Type\":\"String\"},{\"Name\":\"ComponentConfig\",\"Type\":\"String\"},{\"Name\":\"AppType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "SaveFlogoComponents",
            "name": "SaveFlogoComponents",
            "description": "This activity write incoming object to file system",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[BuildFlogoAppFilePath]",
              "accumulate": true
            },
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "%Filename%",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": "=\"ext:app\"==$iteration[value].Data.AppType",
                "Data": {
                  "mapping": {
                    "Input": "=$iteration[value].Data.ComponentConfig"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$iteration[value].Data.FlogoAppFilePath"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Log",
            "name": "Log",
            "description": "Logs a message",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[SaveFlogoComponents]",
              "accumulate": false
            },
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$iteration[value].Filename",
                "addDetails": false
              }
            }
          },
          {
            "id": "BuildFlogoApp",
            "name": "BuildFlogoApp",
            "description": "This activity build docker image from Dockerfile",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[SaveFlogoComponents]",
              "accumulate": false
            },
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "",
                "numOfExecutions": 1,
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"Version\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"}]",
                "SystemEnv": "[{\"Key\":\"FLOGO_HOME\",\"Value\":\"/usr/flogo/home\",\"PerCommand\":\"No\"},{\"Key\":\"PWD\",\"Value\":\"/usr/bin\",\"PerCommand\":\"No\"}]"
              },
              "input": {
                "Asynchronous": "=true",
                "SkipCondition": "=null == $iteration[value].Filename",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "/usr/flogo/home/bin/%FlogoBuilder% build -f %Filename% -docker -tag=%Version%"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$iteration[value].Filename",
                    "Version": "0.1.0",
                    "FlogoBuilder": "=$property[\"System.FlogoBuilder\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Version\":{\"type\":\"string\"},\"FlogoBuilder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\",\"Version\":\"2\",\"FlogoBuilder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"Version\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"}]"
                  },
                  "SystemEnv": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Key\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"PerCommand\":{\"type\":\"string\"}},\"required\":[\"Key\",\"Value\",\"PerCommand\"]}}",
                    "fe_metadata": "[{\"Key\":\"FLOGO_HOME\",\"Value\":\"/usr/flogo/home\",\"PerCommand\":\"No\"},{\"Key\":\"PWD\",\"Value\":\"/usr/bin\",\"PerCommand\":\"No\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildExtApp",
            "name": "BuildExtApp",
            "description": "This activity build docker image from Dockerfile",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[PipelineBuilder].PipelineConfig",
              "accumulate": false
            },
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "",
                "numOfExecutions": 2,
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Version\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=true",
                "SkipCondition": "=\"ext:app\"!=$iteration[value].Type",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_1": "docker push bigoyang/%ID%:%Version%",
                      "Execution_0": "docker build --tag bigoyang/%ID%:%Version% %ProjectFolder%/docker"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "ID": "=string.toLower($iteration[value].Name)",
                    "Version": "0.1.0",
                    "ProjectFolder": "=$activity[ProjectProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Version\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\",\"ID\":\"2\",\"Version\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Version\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::Build General Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "Runner",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AppDescriptor",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:Docker_Container",
      "data": {
        "name": "Docker Container",
        "description": "Build descriptors and images for docker based deployment (include docker-compose and k8s)",
        "links": [
          {
            "id": 1,
            "from": "LogBuildDockerContainer",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "DockerizedService",
            "type": "default"
          },
          {
            "id": 3,
            "from": "DockerizedService",
            "to": "HasAppDescriptor",
            "type": "expression",
            "label": "DockerizedServicetoIsFlogoApp02",
            "value": "$activity[SetSystemProperties].Data.HasAppDescriptor"
          },
          {
            "id": 4,
            "from": "HasAppDescriptor",
            "to": "SaveFlogoAppJson",
            "type": "default"
          },
          {
            "id": 5,
            "from": "SaveFlogoAppJson",
            "to": "BuildFlogoAppAndImage",
            "type": "default"
          },
          {
            "id": 6,
            "from": "BuildFlogoAppAndImage",
            "to": "GenerateF1Descriptor01",
            "type": "default"
          },
          {
            "id": 7,
            "from": "GenerateF1Descriptor01",
            "to": "ReturnBuildFlogoApp",
            "type": "default"
          },
          {
            "id": 8,
            "from": "DockerizedService",
            "to": "NoAppDescriptor",
            "type": "exprOtherwise",
            "label": "DockerizedServicetoNotFlogoApp02"
          },
          {
            "id": 9,
            "from": "NoAppDescriptor",
            "to": "FileUserBuildScript",
            "type": "default"
          },
          {
            "id": 10,
            "from": "FileUserBuildScript",
            "to": "UserDefinedBuild",
            "type": "expression",
            "label": "FileReadertoExecCommand",
            "value": "null!=$activity[FileUserBuildScript].Results \u0026\u0026 0!=array.count($activity[FileUserBuildScript].Results)"
          },
          {
            "id": 11,
            "from": "UserDefinedBuild",
            "to": "ExecUserBuildScript",
            "type": "default"
          },
          {
            "id": 12,
            "from": "ExecUserBuildScript",
            "to": "GenerateF1Descriptor02",
            "type": "default"
          },
          {
            "id": 13,
            "from": "GenerateF1Descriptor02",
            "to": "Return03",
            "type": "default"
          },
          {
            "id": 14,
            "from": "FileUserBuildScript",
            "to": "DefaultBuild",
            "type": "exprOtherwise",
            "label": "FileReadertoExecCommand1"
          },
          {
            "id": 15,
            "from": "DefaultBuild",
            "to": "ExecDefaultBuild",
            "type": "expression",
            "label": "CopyOfCopyOfCopyOfLog1toExecDefaultBuild",
            "value": "null!=$flow.Runner"
          },
          {
            "id": 16,
            "from": "ExecDefaultBuild",
            "to": "GenerateF1Descriptor03",
            "type": "default"
          },
          {
            "id": 17,
            "from": "GenerateF1Descriptor03",
            "to": "Return04",
            "type": "default"
          },
          {
            "id": 18,
            "from": "DefaultBuild",
            "to": "GenerateF1Descriptor04",
            "type": "exprOtherwise",
            "label": "CopyOfCopyOfCopyOfLog1 to ConstructF1Dscrptr04"
          },
          {
            "id": 19,
            "from": "GenerateF1Descriptor04",
            "to": "Return05",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogBuildDockerContainer",
            "name": "LogBuildDockerContainer",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%%% Build Docker Container %%%%%% ID : \", $flow.ID, \", DeployDescriptor : \", coerce.toString($flow.DeployDescriptor), \", SyncBuild : \", coerce.toString($flow.SyncBuild), \", ComponentType : \", $flow.ComponentType, \", Runner : \", $flow.Runner, \", Image : \", $flow.Image, \", AppDescriptor : \", $flow.AppDescriptor)"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"HasAppDescriptor\",\"Type\":\"Boolean\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ID",
                    "ProjectArtifactsFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/artifacts\"",
                    "ProjectDescriptorFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline\"",
                    "ProjectBuildFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/build\"",
                    "DefaultBuildFolderInt": "/home/scripts",
                    "DescriptorFilename": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline/\" + $flow.DeployDescriptor.Name + \".json\"",
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false",
                    "HasAppDescriptor": "=null!=$flow.AppDescriptor \u0026\u0026 \"\"!=$flow.AppDescriptor"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"HasAppDescriptor\":{\"type\":\"boolean\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true,\"HasAppDescriptor\":true,\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"HasAppDescriptor\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true,\"HasAppDescriptor\":true}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"HasAppDescriptor\",\"Type\":\"Boolean\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "DockerizedService",
            "name": "DockerizedService",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=\"################### Component Type = \" + $flow.ComponentType",
                "addDetails": false
              }
            }
          },
          {
            "id": "HasAppDescriptor",
            "name": "HasAppDescriptor",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## docker:has application descriptor ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "SaveFlogoAppJson",
            "name": "SaveFlogoAppJson",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "%Filename%",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": "=false",
                "Data": {
                  "mapping": {
                    "Input": "=$flow.AppDescriptor"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SetSystemProperties].Data.ProjectArtifactsFolderInt + \"/\" + $flow.DeployDescriptor.Name + \".json\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildFlogoAppAndImage",
            "name": "BuildFlogoAppAndImage",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": {
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-exec",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Execution Events Connection",
                  "description": "This connector link exec and execlistener",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/exec",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Connection Name",
                        "description": "Name of the connection"
                      },
                      "value": "DockerBuildEvent"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "Connection description"
                      },
                      "value": ""
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/ModelOps/connector/exec/connector.json",
                  "display": {
                    "description": "This connector link exec and execlistener",
                    "category": "ModelOps",
                    "visible": true,
                    "smallIcon": "exec.png"
                  },
                  "actions": [
                    {
                      "name": "Connect",
                      "display": {
                        "readonly": false,
                        "valid": true,
                        "visible": true
                      }
                    }
                  ],
                  "feature": {},
                  "loopType": "none",
                  "loopSettings": [],
                  "retrySettings": [],
                  "propertyMap": {},
                  "keyfield": "name",
                  "schemaMap": {},
                  "iteratorAccumulate": false,
                  "isValid": true,
                  "lastUpdatedTime": 1613274131095,
                  "createdTime": 1613274131095,
                  "user": "flogo",
                  "subscriptionId": "flogo_sbsc",
                  "id": "9e991180-6e76-11eb-90dc-41e63340b2bf",
                  "connectorName": "DockerBuildEvent",
                  "connectorDescription": " "
                },
                "workingFolder": "$BuildFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"},{\"Name\":\"AppBinFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Runner\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"},{\"Name\":\"ScriptName\",\"Type\":\"String\"},{\"Name\":\"ExecutableName\",\"Type\":\"String\"},{\"Name\":\"BuildImage\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=true",
                "SkipCondition": "=false",
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SaveFlogoAppJson].Filename",
                    "FlogoBuilder": "=$property[\"System.FlogoBuilder\"]",
                    "AppBinFolder": ".",
                    "ServiceName": "=$flow.DeployDescriptor.Name",
                    "ProjectID": "=$flow.ID",
                    "Runner": "flogo",
                    "ImageName": "=$flow.Image",
                    "BuildFolder": "=$activity[SetSystemProperties].Data.DefaultBuildFolderInt",
                    "ScriptName": "build.sh",
                    "ExecutableName": "flogo-engine",
                    "BuildImage": "=f1.ternary(null==$flow.Image||\"\"==$flow.Image, \"no\", \"yes\")"
                  }
                },
                "Executable": {
                  "mapping": {
                    "SystemEnvs": "=$flow.DeployDescriptor.ScriptSystemEnv",
                    "Executions": {
                      "Execution_0": "./$ScriptName$ $ProjectID$ $Runner$ $ImageName$ $FlogoBuilder$ $Filename$ $ServiceName$ $AppBinFolder$ $ExecutableName$ $BuildImage$"
                    }
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"FlogoBuilder\":{\"type\":\"string\"},\"AppBinFolder\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"},\"BuildFolder\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ExecutableName\":{\"type\":\"string\"},\"BuildImage\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\",\"FlogoBuilder\":\"2\",\"AppBinFolder\":\"2\",\"ServiceName\":\"2\",\"ProjectID\":\"2\",\"Runner\":\"2\",\"ImageName\":\"2\",\"BuildFolder\":\"2\",\"ScriptName\":\"2\",\"ExecutableName\":\"2\",\"BuildImage\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"},{\"Name\":\"AppBinFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Runner\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"},{\"Name\":\"ScriptName\",\"Type\":\"String\"},{\"Name\":\"ExecutableName\",\"Type\":\"String\"},{\"Name\":\"BuildImage\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor01",
            "name": "GenerateF1Descriptor01",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": {
                  "mapping": {
                    "Type": "=$flow.DeployDescriptor.Type",
                    "Name": "=$flow.DeployDescriptor.Name",
                    "Descriptor": "=$flow.DeployDescriptor.Descriptor",
                    "Variables": "=$flow.DeployDescriptor.Variables",
                    "ScriptSystemEnv": "=$flow.DeployDescriptor.ScriptSystemEnv",
                    "Properties": "=$flow.DeployDescriptor.Properties"
                  }
                },
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnBuildFlogoApp",
            "name": "ReturnBuildFlogoApp",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "ErrorCode": "=$activity[GenerateF1Descriptor01].ErrorCode",
                  "Success": "=$activity[GenerateF1Descriptor01].Success",
                  "Message": "=$activity[GenerateF1Descriptor01].Message"
                }
              }
            }
          },
          {
            "id": "NoAppDescriptor",
            "name": "NoAppDescriptor",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## docker:not flogo app ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "FileUserBuildScript",
            "name": "FileUserBuildScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SetSystemProperties].Data.ProjectBuildFolderInt + \"/build.sh\""
              },
              "schemas": {
                "output": {
                  "Results": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Content\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Filename\":\"2\",\"Content\":\"2\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "UserDefinedBuild",
            "name": "UserDefinedBuild",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## docker:with customer script ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUserBuildScript",
            "name": "ExecUserBuildScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$BuildFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false!=$activity[SetSystemProperties].Data.SyncBuild",
                "SkipCondition": "=false",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./build.sh",
                      "Execution_1": "./build.sh $ProjectID$ $ServiceName$ $BuildFolder$ TEST"
                    },
                    "SystemEnvs": "=f1.coalesceobject($flow.DeployDescriptor.ScriptSystemEnv, f1.json2object(\"{}\"))"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ID",
                    "ServiceName": "=$flow.DeployDescriptor.Name",
                    "BuildFolder": "=$activity[SetSystemProperties].Data.ProjectBuildFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"BuildFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ServiceName\":\"2\",\"BuildFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor02",
            "name": "GenerateF1Descriptor02",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return03",
            "name": "Return03",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[ExecUserBuildScript].Success",
                  "Message": "=coerce.toString($activity[ExecUserBuildScript].Message)",
                  "ErrorCode": "=$activity[GenerateF1Descriptor02].ErrorCode"
                }
              }
            }
          },
          {
            "id": "DefaultBuild",
            "name": "DefaultBuild",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## docker:with default script ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultBuild",
            "name": "ExecDefaultBuild",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$BuildFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Runner\",\"Type\":\"String\"},{\"Name\":\"DockerUser\",\"Type\":\"String\"},{\"Name\":\"DockerPassword\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false!=$activity[SetSystemProperties].Data.SyncBuild",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_1": "./build-image.sh $ProjectID$ $Runner$ $ImageName$",
                      "Execution_0": "docker login --username=$DockerUser$ --password=$DockerPassword$"
                    },
                    "SystemEnvs": "=f1.coalesceobject($flow.DeployDescriptor.ScriptSystemEnv, f1.json2object(\"{}\"))"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ID",
                    "Runner": "=$flow.Runner",
                    "DockerUser": "bigoyang",
                    "DockerPassword": "43Veryang",
                    "ImageName": "=$flow.Image",
                    "BuildFolder": "=$activity[SetSystemProperties].Data.DefaultBuildFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"DockerUser\":{\"type\":\"string\"},\"DockerPassword\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"},\"BuildFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Runner\":\"2\",\"DockerUser\":\"2\",\"DockerPassword\":\"2\",\"ImageName\":\"2\",\"BuildFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Runner\",\"Type\":\"String\"},{\"Name\":\"DockerUser\",\"Type\":\"String\"},{\"Name\":\"DockerPassword\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor03",
            "name": "GenerateF1Descriptor03",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return04",
            "name": "Return04",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[ExecDefaultBuild].Success",
                  "Message": "=coerce.toString($activity[ExecDefaultBuild].Message)",
                  "ErrorCode": "=$activity[ExecDefaultBuild].ErrorCode"
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor04",
            "name": "GenerateF1Descriptor04",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return05",
            "name": "Return05",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "ErrorCode": "=$activity[GenerateF1Descriptor04].ErrorCode",
                  "Success": "=$activity[GenerateF1Descriptor04].Success",
                  "Message": "=$activity[GenerateF1Descriptor04].Message"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return6",
              "name": "Return6",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::Docker Container - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "Runner",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AppDescriptor",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:AWS_Lambda",
      "data": {
        "name": "AWS Lambda",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "SetSystemProperties",
            "to": "AWSLambdaService",
            "type": "default"
          },
          {
            "id": 2,
            "from": "AWSLambdaService",
            "to": "IsFlogoApp",
            "type": "expression",
            "label": "BuildServiceLogtoSaveFlogoService",
            "value": "null!=$flow.AppDescriptor"
          },
          {
            "id": 3,
            "from": "IsFlogoApp",
            "to": "SaveFlogoService",
            "type": "default"
          },
          {
            "id": 4,
            "from": "SaveFlogoService",
            "to": "BuildFlogoService",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildFlogoService",
            "to": "GenerateF1Descriptor01",
            "type": "default"
          },
          {
            "id": 6,
            "from": "GenerateF1Descriptor01",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 7,
            "from": "AWSLambdaService",
            "to": "NotFlogoApp",
            "type": "exprOtherwise",
            "label": "BuildServiceLog to CopyOfFileReader"
          },
          {
            "id": 8,
            "from": "NotFlogoApp",
            "to": "ExecUserBuildScript01",
            "type": "default"
          },
          {
            "id": 9,
            "from": "ExecUserBuildScript01",
            "to": "GenerateF1Descriptor02",
            "type": "default"
          },
          {
            "id": 10,
            "from": "GenerateF1Descriptor02",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ID",
                    "ProjectArtifactsFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/artifacts\"",
                    "ProjectDescriptorFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline\"",
                    "ProjectBuildFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/build\"",
                    "DefaultBuildFolderInt": "/home/scripts",
                    "DescriptorFilename": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline/descriptor.json\"",
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true,\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "AWSLambdaService",
            "name": "AWSLambdaService",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=\"################### Component Type = \" + $flow.ComponentType",
                "addDetails": false
              }
            }
          },
          {
            "id": "IsFlogoApp",
            "name": "IsFlogoApp",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## aws-lambda:with flogo app ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "SaveFlogoService",
            "name": "SaveFlogoService",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "%Filename%",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": "=false",
                "Data": {
                  "mapping": {
                    "Input": "=$flow.AppDescriptor"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SetSystemProperties].Data.ProjectArtifactsFolderInt + \"/\" + $flow.DeployDescriptor.Name + \".json\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildFlogoService",
            "name": "BuildFlogoService",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "",
                "numOfExecutions": 1,
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"},{\"Name\":\"AppBinFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"Plateform\",\"Type\":\"String\"}]",
                "SystemEnv": "[{\"Key\":\"FLOGO_HOME\",\"Value\":\"/usr/flogo/home\",\"PerCommand\":\"No\"},{\"Key\":\"PWD\",\"Value\":\"/usr/bin\",\"PerCommand\":\"No\"}]"
              },
              "input": {
                "Asynchronous": "=false==$activity[SetSystemProperties].Data.SyncBuild",
                "SkipCondition": "=false",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "/usr/flogo/home/bin/%FlogoBuilder% build -p %Plateform% -f %Filename% -n %ServiceName% -o %AppBinFolder%"
                    },
                    "SystemEnvs": "=$flow.DeployDescriptor.ScriptSystemEnv"
                  }
                },
                "Variables": {
                  "mapping": {
                    "FlogoBuilder": "=$property[\"System.FlogoBuilder\"]",
                    "Filename": "=$activity[SaveFlogoService].Filename",
                    "AppBinFolder": "=$activity[SetSystemProperties].Data.ProjectArtifactsFolderInt + \"/\" + $flow.DeployDescriptor.Name",
                    "ServiceName": "=$flow.DeployDescriptor.Name",
                    "Plateform": "linux/amd64"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"FlogoBuilder\":{\"type\":\"string\"},\"AppBinFolder\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\",\"FlogoBuilder\":\"2\",\"AppBinFolder\":\"2\",\"ServiceName\":\"2\",\"Plateform\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"FlogoBuilder\",\"Type\":\"String\"},{\"Name\":\"AppBinFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"Plateform\",\"Type\":\"String\"}]"
                  },
                  "SystemEnv": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Key\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"PerCommand\":{\"type\":\"string\"}},\"required\":[\"Key\",\"Value\",\"PerCommand\"]}}",
                    "fe_metadata": "[{\"Key\":\"FLOGO_HOME\",\"Value\":\"/usr/flogo/home\",\"PerCommand\":\"No\"},{\"Key\":\"PWD\",\"Value\":\"/usr/bin\",\"PerCommand\":\"No\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor01",
            "name": "GenerateF1Descriptor01",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return01",
            "name": "Return01",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=coerce.toString($activity[BuildFlogoService].Message)",
                  "ErrorCode": "=$activity[BuildFlogoService].ErrorCode",
                  "Success": "=$activity[BuildFlogoService].Success"
                }
              }
            }
          },
          {
            "id": "NotFlogoApp",
            "name": "NotFlogoApp",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "########## aws-lambda:not flogo app ##########",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUserBuildScript01",
            "name": "ExecUserBuildScript01",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ArtifactsFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ArtifactsFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false==$activity[SetSystemProperties].Data.SyncBuild",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./build.sh",
                      "Execution_1": "./build.sh $ProjectID$"
                    },
                    "SystemEnvs": "=f1.coalesceobject($flow.DeployDescriptor.ScriptSystemEnv, f1.json2object(\"{}\"))"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ID",
                    "ServiceName": "=$flow.DeployDescriptor.Name",
                    "ArtifactsFolder": "=$activity[SetSystemProperties].Data.ProjectArtifactsFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"ArtifactsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ServiceName\":\"2\",\"ArtifactsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ArtifactsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor02",
            "name": "GenerateF1Descriptor02",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return02",
            "name": "Return02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "ErrorCode": "=$activity[ExecUserBuildScript01].ErrorCode",
                  "Message": "=coerce.toString($activity[ExecUserBuildScript01].Message)",
                  "Success": "=$activity[ExecUserBuildScript01].Success"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::AWS Lambda - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "Runner",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AppDescriptor",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Build_Service",
      "data": {
        "name": "Build Service",
        "description": "The entry point for general building process",
        "links": [
          {
            "id": 1,
            "from": "LogBuildService",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "ToAWSLambdaFlow",
            "type": "expression",
            "label": "SetSystemPropertiestoSaveFlogoService",
            "value": "\"aws-lambda\" == $flow.DeployDescriptor.Type"
          },
          {
            "id": 3,
            "from": "ToAWSLambdaFlow",
            "to": "ReturnFromAWSLambdaFlow",
            "type": "default"
          },
          {
            "id": 4,
            "from": "SetSystemProperties",
            "to": "ToDockerContainerFlow",
            "type": "expression",
            "label": "SetSystemPropertiestoExecCommand",
            "value": "\"docker\" == $flow.DeployDescriptor.Type"
          },
          {
            "id": 5,
            "from": "ToDockerContainerFlow",
            "to": "ReturnFromDockerContainerFlow",
            "type": "default"
          },
          {
            "id": 6,
            "from": "SetSystemProperties",
            "to": "ToUserDefinedFlow",
            "type": "expression",
            "label": "SetSystemPropertiestoReturnUnknown",
            "value": "\"user-defined\" == $flow.DeployDescriptor.Type"
          },
          {
            "id": 7,
            "from": "ToUserDefinedFlow",
            "to": "ReturnFromUserDefined",
            "type": "default"
          },
          {
            "id": 8,
            "from": "SetSystemProperties",
            "to": "ReturnUnknown",
            "type": "exprOtherwise",
            "label": "SetSystemProperties to ReturnUnknown"
          }
        ],
        "tasks": [
          {
            "id": "LogBuildService",
            "name": "LogBuildService",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%%% Build Service %%%%%% ID : \", $flow.ProjectID, \", SyncBuild : \", coerce.toString($flow.SyncBuild), \"ComponentType : \", $flow.ComponentType, \"Runner : \", $flow.Runner, \", Image : \", $flow.Image, \", AppDescriptor : \", $flow.AppDescriptor, \", DeployDescriptor : \", coerce.toString($flow.DeployDescriptor))"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "ProjectArtifactsFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/artifacts\"",
                    "ProjectDescriptorFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/pipeline\"",
                    "ProjectBuildFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/build\"",
                    "DefaultBuildFolderInt": "/home/scripts",
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"SyncBuild\":true,\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"SyncBuild\":true}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ToAWSLambdaFlow",
            "name": "ToAWSLambdaFlow",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:AWS_Lambda"
              },
              "input": {
                "ID": "=$flow.ProjectID",
                "ComponentType": "=$flow.ComponentType",
                "Runner": "=$flow.Runner",
                "Image": "=$flow.Image",
                "AppDescriptor": "=$flow.AppDescriptor",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "SyncBuild": "=$activity[SetSystemProperties].Data.SyncBuild"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnFromAWSLambdaFlow",
            "name": "ReturnFromAWSLambdaFlow",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToAWSLambdaFlow].Message",
                  "ErrorCode": "=$activity[ToAWSLambdaFlow].ErrorCode",
                  "Success": "=$activity[ToAWSLambdaFlow].Success"
                }
              }
            }
          },
          {
            "id": "ToDockerContainerFlow",
            "name": "ToDockerContainerFlow",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Docker_Container"
              },
              "input": {
                "ID": "=$flow.ProjectID",
                "Image": "=$flow.Image",
                "AppDescriptor": "=$flow.AppDescriptor",
                "ComponentType": "=$flow.ComponentType",
                "Runner": "=$flow.Runner",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "SyncBuild": "=$activity[SetSystemProperties].Data.SyncBuild"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnFromDockerContainerFlow",
            "name": "ReturnFromDockerContainerFlow",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToDockerContainerFlow].Message",
                  "Success": "=$activity[ToDockerContainerFlow].Success",
                  "ErrorCode": "=$activity[ToDockerContainerFlow].ErrorCode"
                }
              }
            }
          },
          {
            "id": "ToUserDefinedFlow",
            "name": "ToUserDefinedFlow",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:UserDefined"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnFromUserDefined",
            "name": "ReturnFromUserDefined",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToUserDefinedFlow].Message",
                  "Success": "=$activity[ToUserDefinedFlow].Success",
                  "ErrorCode": "=$activity[ToUserDefinedFlow].ErrorCode"
                }
              }
            }
          },
          {
            "id": "ReturnUnknown",
            "name": "ReturnUnknown",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=false",
                  "Message": "Unknown project type !!!",
                  "ErrorCode": 100
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return",
              "name": "Return",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::Build Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "Runner",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AppDescriptor",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ProjectID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:UserDefined",
      "data": {
        "name": "UserDefined",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "SetSystemProperties",
            "to": "ExecUserBuildScript01",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ExecUserBuildScript01",
            "to": "GenerateF1Descriptor02",
            "type": "default"
          },
          {
            "id": 3,
            "from": "GenerateF1Descriptor02",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ID",
                    "ProjectArtifactsFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/artifacts\"",
                    "ProjectDescriptorFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline\"",
                    "ProjectBuildFolderInt": "=\"/home/f1/projects/\" + $flow.ID + \"/build\"",
                    "DefaultBuildFolderInt": "/home/scripts",
                    "DescriptorFilename": "=\"/home/f1/projects/\" + $flow.ID + \"/pipeline/descriptor.json\"",
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true,\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"ProjectArtifactsFolderInt\":{\"type\":\"string\"},\"ProjectDescriptorFolderInt\":{\"type\":\"string\"},\"ProjectBuildFolderInt\":{\"type\":\"string\"},\"DefaultBuildFolderInt\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderInt\":\"2\",\"ProjectArtifactsFolderInt\":\"2\",\"ProjectDescriptorFolderInt\":\"2\",\"ProjectBuildFolderInt\":\"2\",\"DefaultBuildFolderInt\":\"2\",\"DescriptorFilename\":\"2\",\"SyncBuild\":true}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectArtifactsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectDescriptorFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DefaultBuildFolderInt\",\"Type\":\"String\"},{\"Name\":\"DescriptorFilename\",\"Type\":\"String\"},{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ExecUserBuildScript01",
            "name": "ExecUserBuildScript01",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ArtifactsFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ArtifactsFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false==$activity[SetSystemProperties].Data.SyncBuild",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./build.sh",
                      "Execution_1": "./build.sh $ProjectID$"
                    },
                    "SystemEnvs": "=f1.coalesceobject($flow.DeployDescriptor.ScriptSystemEnv, f1.json2object(\"{}\"))"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ID",
                    "ServiceName": "=$flow.DeployDescriptor.Name",
                    "ArtifactsFolder": "=$activity[SetSystemProperties].Data.ProjectArtifactsFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"ArtifactsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ServiceName\":\"2\",\"ArtifactsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ArtifactsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateF1Descriptor02",
            "name": "GenerateF1Descriptor02",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Generate_F1_Descriptor"
              },
              "input": {
                "ID": "=$flow.ID",
                "DeployDescriptor": "=$flow.DeployDescriptor",
                "DescriptorFilename": "=$activity[SetSystemProperties].Data.DescriptorFilename"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"DescriptorFilename\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return02",
            "name": "Return02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "ErrorCode": "=$activity[ExecUserBuildScript01].ErrorCode",
                  "Message": "=coerce.toString($activity[ExecUserBuildScript01].Message)",
                  "Success": "=$activity[ExecUserBuildScript01].Success"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"Builder::AWS Lambda - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "Runner",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AppDescriptor",
              "type": "string"
            },
            {
              "name": "DeployDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ID\":{\"type\":\"string\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:Ping",
      "data": {
        "name": "Ping",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "BuildServiceInfo",
            "to": "HTTPClient",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "BuildServiceInfo",
            "name": "BuildServiceInfo",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": {
                  "filename": "ServiceLoaderRegist.json",
                  "content": "data:application/json;base64,ewogICAgIklEIjogInByb2plY3RtZ3ItMTIzLjQ4Ni43ODkiLAogICAgIlR5cGUiOiAicHJvamVjdG1nciIsCiAgICAiVVJMIjogImh0dHBzOi8vYXBpLmdpdGh1Yi5jb20vdXNlcnMvU3RldmVOWS1UaWJjby9yZXBvcyIKfQ=="
                },
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "URL": "=\"http://\" + $property[\"System.ServiceLocatorIP\"] + \":10083/f1/builder\"",
                    "Type": "builder",
                    "ID": "=\"builder_\" + $property[\"System.ServiceLocatorIP\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"\",\"Type\":\"\",\"URL\":\"\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "HTTPClient",
            "name": "HTTPClient",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:10080/f1/locator/register/%ServiceType%",
                "Body": "=$activity[BuildServiceInfo].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "ServiceType": "builder"
                  }
                }
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Now",
              "type": "string"
            }
          ],
          "output": [],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Now\":{\"type\":\"string\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:List_Runner",
      "data": {
        "name": "List Runner",
        "description": "",
        "links": [],
        "tasks": [
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn"
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Command",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Runner",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Decription\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Runner\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Decription\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Runner",
      "data": {
        "name": "Get Runner",
        "description": "",
        "links": [],
        "tasks": [
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn"
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "ID",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Push_And_Build_Project",
      "data": {
        "name": "Push And Build Project",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ExecCommand",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "ExecCommand",
            "name": "ExecCommand",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "numOfExecutions": 1,
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"CompressedFile\",\"Type\":\"String\"},{\"Name\":\"ExtractTo\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": false,
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "tar -zxf %CompressedFile% -C %ExtractTo%"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "CompressedFile": "=$flow.FilePath+\"\"+$flow.Filename",
                    "ExtractTo": "=$property[\"System.BaseFolder\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"CompressedFile\":{\"type\":\"string\"},\"ExtractTo\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"CompressedFile\":\"2\",\"ExtractTo\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"CompressedFile\",\"Type\":\"String\"},{\"Name\":\"ExtractTo\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "OutputData": "=$activity[ExecCommand].Message"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Filename",
              "type": "string"
            },
            {
              "name": "FilePath",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "OutputData",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"FilePath\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"OutputData\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Build_Event",
      "data": {
        "name": "Build Event",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "ObjectMaker",
            "type": "expression",
            "label": "LogMessagetoObjectMaker",
            "value": "null!=$flow.Event.SystemEnvironment.RSVPUrl"
          },
          {
            "id": 2,
            "from": "ObjectMaker",
            "to": "ObjectSerializer",
            "type": "expression",
            "label": "ObjectMaker to ObjectSerializer",
            "value": "$flow.Event.Successful"
          },
          {
            "id": 3,
            "from": "ObjectSerializer",
            "to": "HTTPClient",
            "type": "default"
          },
          {
            "id": 4,
            "from": "ObjectMaker",
            "to": "LogError",
            "type": "exprOtherwise",
            "label": "ObjectMakertoLogMessage1"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=\"######## BUILD EVENT ######### \" + coerce.toString($flow.Event)"
              }
            }
          },
          {
            "id": "ObjectMaker",
            "name": "ObjectMaker",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Successful\" : true,\n    \"ErrorMsg\" : \"string\"\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "Successful": "=$flow.Event.Successful",
                    "ErrorMsg": "=$flow.Event.ErrorMsg"
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Successful\":true,\"ErrorMsg\":\"string\"}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Successful\":true,\"ErrorMsg\":\"string\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"Successful\" : true,\n    \"ErrorMsg\" : \"string\"\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ObjectSerializer",
            "name": "ObjectSerializer",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[ObjectMaker].ObjectOut"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "HTTPClient",
            "name": "HTTPClient",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "30000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$flow.Event.SystemEnvironment.RSVPUrl",
                "Body": "=$activity[ObjectSerializer].SerializedString",
                "SkipCondition": false
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "LogError",
            "name": "LogError",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=\"######## Building Pipeline Fail!! ######### Error message : \" + $flow.Event.ErrorMsg"
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Event",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"SystemEnvironment\":{\"type\":\"object\",\"properties\":{\"RSVPUrl\":{\"type\":\"string\"}}},\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}"
              }
            }
          ],
          "output": [],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Event\":{\"type\":\"object\",\"properties\":{\"SystemEnvironment\":{\"type\":\"object\",\"properties\":{\"RSVPUrl\":{\"type\":\"string\"}}},\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Upload_Resource",
      "data": {
        "name": "Upload Resource",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogResourceInfo",
            "to": "FileWriter",
            "type": "expression",
            "label": "LogResourceInfotoFileWriter",
            "value": "\"GO_LIB\"==$flow.Type"
          },
          {
            "id": 2,
            "from": "FileWriter",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogResourceInfo",
            "name": "LogResourceInfo",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%%% Upload Resource %%%%%% Type : \", $flow.Type, \", ID : \", $flow.ID)"
              }
            }
          },
          {
            "id": "FileWriter",
            "name": "FileWriter",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "/usr/wi/ext/src/$Path$/$ID$.go",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Path\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$flow.Resource"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ID": "=$flow.ID",
                    "Path": "=$flow.Path"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Path\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"Path\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Path\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"Path\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Path\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "=string.concat(\"Resource \", $activity[FileWriter].Filename, \" saved!\")",
                  "ErrorCode": 100
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "Message": "=string.concat(\"Resource \", $flow.Path, \" not saved : \", $error.message)",
                    "ErrorCode": 400
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "Type",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Path",
              "type": "string"
            },
            {
              "name": "Resource",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Path\":{\"type\":\"string\"},\"Resource\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}
