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
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/project-flogo/contrib/function/string",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/project-flogo/legacybridge",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/flow/activity/subflow",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer"
  ],
  "properties": [
    {
      "name": "System.ServiceLocatorIP",
      "type": "string",
      "value": "servicelocator"
    },
    {
      "name": "System.BaseFolder",
      "type": "string",
      "value": "BaseFolder not set ! "
    },
    {
      "name": "System.BaseFolderExt",
      "type": "string",
      "value": "BaseFolderExt not set ! "
    },
    {
      "name": "System.Network",
      "type": "string",
      "value": "Network not set !"
    },
    {
      "name": "System.ServiceLocatorPort",
      "type": "string",
      "value": "10080"
    },
    {
      "name": "System.ProjectBaseFolderInt",
      "type": "string",
      "value": "/home/f1/projects"
    },
    {
      "name": "System.Ping_Interval",
      "type": "float64",
      "value": 10
    },
    {
      "name": "System.AgentID",
      "type": "string",
      "value": "AgentID not set !"
    }
  ],
  "triggers": [
    {
      "id": "TimerTrigger",
      "ref": "#timer",
      "settings": {},
      "handlers": [
        {
          "name": "PingLightCrane",
          "settings": {
            "Interval Unit": "Second",
            "Repeating": true,
            "Start Date": "",
            "Time Interval": "=$property[\"System.Ping_Interval\"]"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:PingLightCrane"
              },
              "input": {
                "Now": ""
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:Envoke_Script",
      "data": {
        "name": "Envoke Script",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogEnvokeScript",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "LogMessage",
            "type": "default"
          },
          {
            "id": 3,
            "from": "LogMessage",
            "to": "DeployWithCustomerScript",
            "type": "default"
          },
          {
            "id": 4,
            "from": "DeployWithCustomerScript",
            "to": "Return_EnvokeScript",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogEnvokeScript",
            "name": "LogEnvokeScript",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"######## Envoke Script ######## SyncDeploy : \", coerce.toString($flow.SyncDeploy), \", ScriptName : \", $flow.ScriptName, \", Namesapce : \", $flow.Namespace, \", DeployFolder : \", $flow.DeployFolder, \", ProjectID : \", $flow.ProjectID, \", DeployType : \", $flow.DeployType, \", ServiceName : \", $flow.ServiceName, \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
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
                "MappingFields": "[{\"Name\":\"SyncDeply\",\"Type\":\"Boolean\"},{\"Name\":\"ScriptSystemEnv\",\"Type\":\"Object\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SkipCondition": "=false",
                    "SyncDeply": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)",
                    "ScriptSystemEnv": "=f1.coalesceobject($flow.ScriptSystemEnv, f1.json2object('{}'))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeply\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"SyncDeply\":true,\"ScriptSystemEnv\":{},\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeply\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"SyncDeply\":true,\"ScriptSystemEnv\":{}}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"SyncDeply\",\"Type\":\"Boolean\"},{\"Name\":\"ScriptSystemEnv\",\"Type\":\"Object\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=\"@@@@@@@@@@@@@@@@@@ : \" + coerce.toString($activity[SetSystemProperties].Data.ScriptSystemEnv)"
              }
            }
          },
          {
            "id": "DeployWithCustomerScript",
            "name": "DeployWithCustomerScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$WorkFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"DeployType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ScriptName\",\"Type\":\"String\"},{\"Name\":\"InstanceID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false==$activity[SetSystemProperties].Data.SyncDeply",
                "SkipCondition": "=false",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./$ScriptName$",
                      "Execution_1": "./$ScriptName$ $Namespace$ $DeployType$ $ServiceName$ $ProjectID$ $InstanceID$"
                    },
                    "SystemEnvs": "=$activity[SetSystemProperties].Data.ScriptSystemEnv"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "Namespace": "=$flow.Namespace",
                    "WorkFolder": "=$flow.DeployFolder",
                    "DeployType": "=$flow.DeployType",
                    "ServiceName": "=$flow.ServiceName",
                    "ScriptName": "=$flow.ScriptName",
                    "InstanceID": "=$flow.InstanceID"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"WorkFolder\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"WorkFolder\":\"2\",\"Namespace\":\"2\",\"DeployType\":\"2\",\"ServiceName\":\"2\",\"ProjectID\":\"2\",\"ScriptName\":\"2\",\"InstanceID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"DeployType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ScriptName\",\"Type\":\"String\"},{\"Name\":\"InstanceID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_EnvokeScript",
            "name": "Return_EnvokeScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DeployWithCustomerScript].Success",
                  "ErrorCode": "=$activity[DeployWithCustomerScript].ErrorCode",
                  "Message": "=coerce.toString($activity[DeployWithCustomerScript].Message)"
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
                    "Message": "=$error.message"
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
              "name": "SyncDeploy",
              "type": "boolean"
            },
            {
              "name": "Namespace",
              "type": "string"
            },
            {
              "name": "DeployFolder",
              "type": "string"
            },
            {
              "name": "DeployType",
              "type": "string"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "ServiceName",
              "type": "string"
            },
            {
              "name": "InstanceID",
              "type": "string"
            },
            {
              "name": "ScriptName",
              "type": "string"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:Unregister",
      "data": {
        "name": "Unregister",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "SetSystemProperties",
            "to": "DirectEnvokeUndeployScript",
            "type": "default"
          },
          {
            "id": 2,
            "from": "DirectEnvokeUndeployScript",
            "to": "BuildUnregisterRequestBody",
            "type": "default"
          },
          {
            "id": 3,
            "from": "BuildUnregisterRequestBody",
            "to": "Unregister",
            "type": "expression",
            "label": "BuildUnregisterRequestBodytoUnregister",
            "value": "null!=$activity[SetSystemProperties].Data.F1ServiceType"
          },
          {
            "id": 4,
            "from": "Unregister",
            "to": "Return_AfterUnregister",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildUnregisterRequestBody",
            "to": "Return_DeirectUndeploy",
            "type": "exprOtherwise",
            "label": "BuildUnregisterRequestBody to CopyOfReturn_DeirectUndeploy"
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
                "MappingFields": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"SyncUndeploy\",\"Type\":\"Boolean\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"F1ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "DeployFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/\"",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "ProjectFolderExt": "=$property[\"System.BaseFolderExt\"] + \"/projects/\"",
                    "SkipCondition": "=false",
                    "SyncUndeploy": "=f1.ternary(null!=$flow.SyncUndeploy, $flow.SyncUndeploy, false)",
                    "ServiceName": "Not defined!",
                    "F1ServiceType": "=f1.getsubobject($flow.ScriptSystemEnv, \"F1ServiceType\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"SyncUndeploy\":{\"type\":\"boolean\"},\"ServiceName\":{\"type\":\"string\"},\"F1ServiceType\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"ProjectFolderExt\":\"2\",\"SyncUndeploy\":true,\"ServiceName\":\"2\",\"F1ServiceType\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"SyncUndeploy\":{\"type\":\"boolean\"},\"ServiceName\":{\"type\":\"string\"},\"F1ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"ProjectFolderExt\":\"2\",\"SyncUndeploy\":true,\"ServiceName\":\"2\",\"F1ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"SyncUndeploy\",\"Type\":\"Boolean\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"F1ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "DirectEnvokeUndeployScript",
            "name": "DirectEnvokeUndeployScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncUndeploy",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                "ProjectID": "=$flow.ProjectID",
                "DeployType": "user-defined",
                "ServiceName": "=$activity[SetSystemProperties].Data.ServiceName",
                "ScriptName": "undeploy.sh",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
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
            "id": "BuildUnregisterRequestBody",
            "name": "BuildUnregisterRequestBody",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\" : \"\"\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "ID": "=$flow.ProjectID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"ID\" : \"\"\n}"
                  }
                }
              }
            }
          },
          {
            "id": "Unregister",
            "name": "Unregister",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"F1ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://$ServiceLocatorIP$:10080/mops/locator/unregister/$F1ServiceType$",
                "Body": "=$activity[BuildUnregisterRequestBody].JSONString",
                "SkipCondition": "=false",
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "F1ServiceType": "=$activity[SetSystemProperties].Data.F1ServiceType"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"F1ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"F1ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"F1ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_AfterUnregister",
            "name": "Return_AfterUnregister",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[Unregister].Success",
                  "Message": "=$activity[Unregister].Data",
                  "ErrorCode": "=$activity[Unregister].ErrorCode"
                }
              }
            }
          },
          {
            "id": "Return_DeirectUndeploy",
            "name": "Return_DeirectUndeploy",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DirectEnvokeUndeployScript].Success",
                  "Message": "=$activity[DirectEnvokeUndeployScript].Message",
                  "ErrorCode": "=$activity[DirectEnvokeUndeployScript].ErrorCode"
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
                    "Message": "=string.concat(\"Deployer::Unregister - \", $error.message)"
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
              "name": "SyncUndeploy",
              "type": "boolean"
            },
            {
              "name": "Method",
              "type": "string"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncUndeploy\":{\"type\":\"boolean\"},\"Method\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:PingLightCrane",
      "data": {
        "name": "PingLightCrane",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "SetSystemProperties",
            "to": "BuildServiceInfo",
            "type": "default"
          },
          {
            "id": 2,
            "from": "BuildServiceInfo",
            "to": "PingServiceLocator",
            "type": "default"
          },
          {
            "id": 3,
            "from": "PingServiceLocator",
            "to": "QueryDeployment",
            "type": "expression",
            "label": "HTTPClienttoHTTPClient1",
            "value": "$activity[PingServiceLocator].Success"
          },
          {
            "id": 4,
            "from": "QueryDeployment",
            "to": "ExtractDeployment",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ExtractDeployment",
            "to": "LogEnvokeScript",
            "type": "expression",
            "label": "ExtractDeploymenttoStartaSubFlow",
            "value": "false == f1.isempty($activity[ExtractDeployment].Data)"
          },
          {
            "id": 6,
            "from": "LogEnvokeScript",
            "to": "DownloadResources",
            "type": "expression",
            "label": "LogEnvokeScript to DownloadResources",
            "value": "\"deploy.sh\"==$activity[ExtractDeployment].Data.Variables.ScriptName"
          },
          {
            "id": 7,
            "from": "DownloadResources",
            "to": "EnvokeDeployScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "LogEnvokeScript",
            "to": "EnvokeUndeployScript",
            "type": "exprOtherwise",
            "label": "LogEnvokeScript to EnvokeUndeployScript"
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
                "MappingFields": "[{\"Name\":\"AgentID\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "AgentID": "=\"agent_\" + $property[\"System.AgentID\"]",
                    "ServiceType": "agent"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"AgentID\":\"2\",\"ServiceType\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"AgentID\":\"2\",\"ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"AgentID\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildServiceInfo",
            "name": "BuildServiceInfo",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "ID": "=$activity[SetSystemProperties].Data.AgentID",
                    "Type": "=$activity[SetSystemProperties].Data.ServiceType",
                    "URL": "=\"http://\" + $property[\"System.ServiceLocatorIP\"] + \":10082/f1/agent\""
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
            "id": "PingServiceLocator",
            "name": "PingServiceLocator",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "500",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:%ServiceLocatorPort%/f1/locator/register/%ServiceType%",
                "Body": "=$activity[BuildServiceInfo].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "ServiceType": "=$activity[SetSystemProperties].Data.ServiceType",
                    "ServiceLocatorPort": "=$property[\"System.ServiceLocatorPort\"]"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"ServiceLocatorPort\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"ServiceType\":\"2\",\"ServiceLocatorPort\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "QueryDeployment",
            "name": "QueryDeployment",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"},{\"Name\":\"AgentID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:%ServiceLocatorPort%/f1/deployer/deployment/%AgentID%",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "ServiceLocatorPort": "=$property[\"System.ServiceLocatorPort\"]",
                    "AgentID": "=$activity[SetSystemProperties].Data.AgentID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"ServiceLocatorPort\":{\"type\":\"string\"},\"AgentID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"ServiceLocatorPort\":\"2\",\"AgentID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"},{\"Name\":\"AgentID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ExtractDeployment",
            "name": "ExtractDeployment",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"AgentID\" : \"\",\n    \"Variables\" : {\n        \"Namespace\" : \"\",\n        \"DeployType\" : \"\",\n        \"ProjectID\" : \"\",\n        \"ServiceName\" : \"\",\n        \"InstanceID\" : \"\",\n        \"ScriptName\" : \"\"\n    },\n    \"SystemEnv\" : {},\n    \"Resources\" : {\n        \"Path\" : \"\",\n        \"File\" : \"\"\n    }\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[QueryDeployment].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}}}",
                    "fe_metadata": "{\"AgentID\":\"\",\"Variables\":{\"Namespace\":\"\",\"DeployType\":\"\",\"ProjectID\":\"\",\"ServiceName\":\"\",\"InstanceID\":\"\",\"ScriptName\":\"\"},\"SystemEnv\":{},\"Resources\":{\"Path\":\"\",\"File\":\"\"}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}}}",
                    "fe_metadata": "{\n    \"AgentID\" : \"\",\n    \"Variables\" : {\n        \"Namespace\" : \"\",\n        \"DeployType\" : \"\",\n        \"ProjectID\" : \"\",\n        \"ServiceName\" : \"\",\n        \"InstanceID\" : \"\",\n        \"ScriptName\" : \"\"\n    },\n    \"SystemEnv\" : {},\n    \"Resources\" : {\n        \"Path\" : \"\",\n        \"File\" : \"\"\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "LogEnvokeScript",
            "name": "LogEnvokeScript",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Envoke Script ^^^^^^ Deployment Data : \", coerce.toString($activity[ExtractDeployment].Data))"
              }
            }
          },
          {
            "id": "DownloadResources",
            "name": "DownloadResources",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "%WorkFolder%",
                "numOfExecutions": 2,
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"},{\"Name\":\"Path\",\"Type\":\"String\"},{\"Name\":\"File\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": false,
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "WorkFolder": "=$property[\"System.ProjectBaseFolderInt\"]",
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "ServiceLocatorPort": "=$property[\"System.ServiceLocatorPort\"]",
                    "Path": "=$activity[ExtractDeployment].Data.Resources.Path",
                    "File": "=$activity[ExtractDeployment].Data.Resources.File"
                  }
                },
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "curl http://%ServiceLocatorIP%:%ServiceLocatorPort%/%Path%/%File% --output %File%",
                      "Execution_1": "unzip -o %File%"
                    }
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"WorkFolder\":{\"type\":\"string\"},\"ServiceLocatorIP\":{\"type\":\"string\"},\"ServiceLocatorPort\":{\"type\":\"string\"},\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"WorkFolder\":\"2\",\"ServiceLocatorIP\":\"2\",\"ServiceLocatorPort\":\"2\",\"Path\":\"2\",\"File\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorPort\",\"Type\":\"String\"},{\"Name\":\"Path\",\"Type\":\"String\"},{\"Name\":\"File\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "EnvokeDeployScript",
            "name": "EnvokeDeployScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "Namespace": "=$activity[ExtractDeployment].Data.Variables.Namespace",
                "DeployFolder": "=string.concat($property[\"System.ProjectBaseFolderInt\"], \"/\", $activity[ExtractDeployment].Data.Variables.ProjectID, \"/deploy/\")",
                "DeployType": "=$activity[ExtractDeployment].Data.Variables.DeployType",
                "ProjectID": "=$activity[ExtractDeployment].Data.Variables.ProjectID",
                "ServiceName": "=$activity[ExtractDeployment].Data.Variables.ServiceName",
                "InstanceID": "=$activity[ExtractDeployment].Data.Variables.InstanceID",
                "ScriptName": "=$activity[ExtractDeployment].Data.Variables.ScriptName",
                "ScriptSystemEnv": "=$activity[ExtractDeployment].Data.SystemEnv"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
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
            "id": "EnvokeUndeployScript",
            "name": "EnvokeUndeployScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "Namespace": "=$activity[ExtractDeployment].Data.Variables.Namespace",
                "DeployFolder": "=string.concat($property[\"System.ProjectBaseFolderInt\"], \"/\", $activity[ExtractDeployment].Data.Variables.ProjectID, \"/deploy/\")",
                "DeployType": "=$activity[ExtractDeployment].Data.Variables.DeployType",
                "ProjectID": "=$activity[ExtractDeployment].Data.Variables.ProjectID",
                "ServiceName": "=$activity[ExtractDeployment].Data.Variables.ServiceName",
                "InstanceID": "=$activity[ExtractDeployment].Data.Variables.InstanceID",
                "ScriptName": "=$activity[ExtractDeployment].Data.Variables.ScriptName",
                "ScriptSystemEnv": "=$activity[ExtractDeployment].Data.SystemEnv"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
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
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}
