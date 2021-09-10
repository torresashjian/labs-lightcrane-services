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
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/project-flogo/contrib/function/array",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/trigger/fileserver",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filewriter",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/properties2object",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/subflow",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/textreplacer",
    "github.com/project-flogo/contrib/function/string",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectmaker",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tablemutate",
    "github.com/project-flogo/legacybridge"
  ],
  "properties": [
    {
      "name": "System.ServiceLocatorIP",
      "type": "string",
      "value": "ServiceLocatorIP not set !"
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
      "name": "System.ExternalEndpointIP",
      "type": "string",
      "value": "ExternalEndpointIP not set !"
    },
    {
      "name": "System.Network",
      "type": "string",
      "value": "Network not set !"
    },
    {
      "name": "System.ServiceLocatorExternal",
      "type": "string",
      "value": "ServiceLocatorExternal not set!"
    },
    {
      "name": "System.ProjectBaseFolderInt",
      "type": "string",
      "value": "/home/f1/projects"
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
        "port": 10082
      },
      "handlers": [
        {
          "name": "Deploy",
          "settings": {
            "method": "POST",
            "path": "/f1/deployer/deploy/:id/:name/:instance"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Deploy"
              },
              "input": {
                "Action": "=$.content.Action",
                "Instance": "=$.pathParams.instance",
                "Method": "=$.content.Method",
                "Name": "=$.pathParams.name",
                "NoF1Descriptor": "=$.content.NoF1Descriptor",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "SyncDeploy": "=$.content.SyncDeploy",
                "UserDefinedParameters": "=$.content.UserDefinedParameters"
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
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"Action\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"Action\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"},{\"parameterName\":\"instance\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"instance\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_Endpoint",
          "settings": {
            "method": "GET",
            "path": "/f1/deployer/endpoint/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Endpoint"
              },
              "input": {
                "ProjectID": "=$.pathParams.id"
              },
              "output": {
                "code": 200,
                "data": "=$.Response"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Deploy_F1_Descriptor",
          "settings": {
            "method": "POST",
            "path": "/f1/deployer/deployF1Descriptor/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Deploy_F1_Descriptor"
              },
              "input": {
                "DeployFolder": "=$.content.DeployFolder",
                "F1Descriptor": "=$.content.F1Descriptor",
                "ProjectFolderInt": "=$.content.ProjectFolderInt",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "SyncDeploy": "=$.content.SyncDeploy",
                "UseDefaultScript": "=$.content.UseDefaultScript"
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
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Undeploy",
          "settings": {
            "method": "POST",
            "path": "/f1/deployer/undeploy/:id/:name/:instance"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Undeploy"
              },
              "input": {
                "Instance": "=$.pathParams.instance",
                "Method": "=$.content.Method",
                "Name": "=$.pathParams.name",
                "NoF1Descriptor": "=$.content.NoF1Descriptor",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "SyncDeploy": "=$.content.SyncDeploy",
                "UserDefinedParameters": "=$.content.UserDefinedParameters"
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
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"},{\"parameterName\":\"instance\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"instance\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "AgentServer",
          "settings": {
            "method": "POST",
            "path": "/f1/deployer/deployment/:AgentID"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:AgentServer"
              },
              "input": {
                "AgentID": "=$.pathParams.AgentID",
                "DeployedComponents": "=$.content.DeployedComponents",
                "NetworkAddress": "=$.content.NetworkAddress"
              },
              "output": {
                "code": 200,
                "data": "=$.Data"
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"NetworkAddress\": {\n            \"type\": \"string\"\n        },\n        \"DeployedComponents\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"ID\": {\n                        \"type\": \"string\"\n                    },\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"NetworkAddress\": {\n            \"type\": \"string\"\n        },\n        \"DeployedComponents\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"ID\": {\n                        \"type\": \"string\"\n                    },\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"AgentID\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"AgentID\" : \"\",\n    \"Variables\" : {\n        \"Namespace\" : \"\",\n        \"DeployFolder\" : \"\",\n        \"DeployType\" : \"\",\n        \"ProjectID\" : \"\",\n        \"ServiceName\" : \"\",\n        \"InstanceID\" : \"\",\n        \"ScriptName\" : \"\"\n    },\n    \"SystemEnv\" : {}\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
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
      "id": "HTTPfileserver1",
      "ref": "#fileserver",
      "settings": {
        "Port": "10084"
      },
      "handlers": [
        {
          "name": "Resources",
          "settings": {
            "Dir": "/home/f1/projects/",
            "URLPath": "/f1/deployer/resources/"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Resources"
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:Create_Deploy_Descriptor",
      "data": {
        "name": "Create Deploy Descriptor",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogCreateDeployDescriptor",
            "to": "BldDeployDescObj",
            "type": "default"
          },
          {
            "id": 2,
            "from": "BldDeployDescObj",
            "to": "GenerateYamlDeployDesc",
            "type": "default"
          },
          {
            "id": 3,
            "from": "GenerateYamlDeployDesc",
            "to": "WriteExtDeployDesc",
            "type": "default"
          },
          {
            "id": 4,
            "from": "WriteExtDeployDesc",
            "to": "ReturnCreateDeployDescriptor",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogCreateDeployDescriptor",
            "name": "LogCreateDeployDescriptor",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"###### Create Deploy Descriptor ###### ProjectFolder : \", $flow.ProjectFolder, \", ProjectID : \", $flow.ProjectID, \", Service : \", coerce.toString($flow.Service))"
              }
            }
          },
          {
            "id": "BldDeployDescObj",
            "name": "BldDeployDescObj",
            "description": "This activity convert f1 key/value to golang object format",
            "activity": {
              "ref": "#properties2object",
              "settings": {
                "defOverrideProperties": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "",
                "PassThrough": "[{\"FieldName\":\"Descriptor\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Name\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Type\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"}]"
              },
              "input": {
                "Properties": "=$flow.Service.Properties",
                "Variables": "=$flow.Service.Variables",
                "PassThroughData": {
                  "mapping": {
                    "Descriptor": "=$flow.Service.Descriptor",
                    "Name": "=$flow.Service.Name",
                    "Type": "=$flow.Service.Type"
                  }
                }
              },
              "schemas": {
                "input": {
                  "OverrideProperties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Type\":\"2\",\"Name\":\"2\",\"Value\":\"2\"}]"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  },
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Descriptor\":\"2\",\"Name\":\"2\",\"Type\":\"2\"}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Descriptor\":\"2\",\"Name\":\"2\",\"Type\":\"2\"}"
                  }
                },
                "settings": {
                  "PassThrough": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"FieldName\":{\"type\":\"string\"},\"Default\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Optional\":{\"type\":\"string\"}},\"required\":[\"FieldName\",\"Default\",\"Type\",\"Optional\"]}}",
                    "fe_metadata": "[{\"FieldName\":\"Descriptor\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Name\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Type\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GenerateYamlDeployDesc",
            "name": "GenerateYamlDeployDesc",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "yaml",
                "PassThrough": "[{\"FieldName\":\"Descriptor\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Name\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Type\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"}]"
              },
              "input": {
                "PassThroughData": {
                  "mapping": {
                    "Descriptor": "=$activity[BldDeployDescObj].PassThroughDataOut.Descriptor",
                    "Name": "=$activity[BldDeployDescObj].PassThroughDataOut.Name",
                    "Type": "=$activity[BldDeployDescObj].PassThroughDataOut.Type"
                  }
                },
                "Data": "=$activity[BldDeployDescObj].DataObject"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Descriptor\":\"2\",\"Name\":\"2\",\"Type\":\"2\"}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Descriptor\":\"2\",\"Name\":\"2\",\"Type\":\"2\"}"
                  }
                },
                "settings": {
                  "PassThrough": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"FieldName\":{\"type\":\"string\"},\"Default\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Optional\":{\"type\":\"string\"}},\"required\":[\"FieldName\",\"Default\",\"Type\",\"Optional\"]}}",
                    "fe_metadata": "[{\"FieldName\":\"Descriptor\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Name\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"},{\"FieldName\":\"Type\",\"Default\":\"\",\"Type\":\"String\",\"Optional\":\"no\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "WriteExtDeployDesc",
            "name": "WriteExtDeployDesc",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$Filename$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": "=false",
                "Variables": {
                  "mapping": {
                    "Filename": "=$flow.ProjectFolder + \"/artifacts/\" + $activity[GenerateYamlDeployDesc].PassThroughDataOut.Descriptor",
                    "Type": "=$activity[GenerateYamlDeployDesc].PassThroughDataOut.Type",
                    "Name": "=$activity[GenerateYamlDeployDesc].PassThroughDataOut.Name"
                  }
                },
                "Data": {
                  "mapping": {
                    "Input": "=$activity[GenerateYamlDeployDesc].SerializedString"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\",\"Type\":\"2\",\"Name\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\",\"Type\":\"2\",\"Name\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnCreateDeployDescriptor",
            "name": "ReturnCreateDeployDescriptor",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Filename": "=$activity[WriteExtDeployDesc].VariablesOut.Filename",
                  "Type": "=$activity[WriteExtDeployDesc].VariablesOut.Type",
                  "Name": "=$activity[WriteExtDeployDesc].VariablesOut.Name"
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
              "name": "ProjectFolder",
              "type": "string"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Service",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Filename",
              "type": "string"
            },
            {
              "name": "Type",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Service\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:Envoke_Script",
      "data": {
        "name": "Envoke Script",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogEnvokeScript",
            "to": "BuildVariables",
            "type": "default"
          },
          {
            "id": 2,
            "from": "BuildVariables",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "GetAgent",
            "type": "default"
          },
          {
            "id": 4,
            "from": "GetAgent",
            "to": "ExtractAgent",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ExtractAgent",
            "to": "LogBeforeEnvoke",
            "type": "default"
          },
          {
            "id": 6,
            "from": "LogBeforeEnvoke",
            "to": "DeployWithCustomerScript",
            "type": "expression",
            "label": "LogMessagetoDeployWithCustomerScript",
            "value": "f1.isempty($activity[ExtractAgent].Data)"
          },
          {
            "id": 7,
            "from": "DeployWithCustomerScript",
            "to": "Return_EnvokeScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "LogBeforeEnvoke",
            "to": "BuildResources",
            "type": "exprOtherwise",
            "label": "LogMessagetoTableMutate"
          },
          {
            "id": 9,
            "from": "BuildResources",
            "to": "TableMutate",
            "type": "default"
          },
          {
            "id": 10,
            "from": "TableMutate",
            "to": "ZipProjectFolder",
            "type": "expression",
            "label": "TableMutatetoZipProjectFolder",
            "value": "\"deploy.sh\"==$activity[BuildVariables].ObjectOut.ScriptName"
          },
          {
            "id": 11,
            "from": "ZipProjectFolder",
            "to": "Return_EnvokeDeployScriptRemote",
            "type": "default"
          },
          {
            "id": 12,
            "from": "TableMutate",
            "to": "Return_EnvokeUndeployScriptRemote",
            "type": "exprOtherwise",
            "label": "TableMutatetoReturn_EnvokeUndeployScriptRemote"
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
            "id": "BuildVariables",
            "name": "BuildVariables",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Namespace\" : \"\",\n    \"DeployType\" : \"\",\n    \"ProjectID\" : \"\",\n    \"ServiceName\" : \"\",\n    \"InstanceID\" : \"\",\n    \"ScriptName\" : \"\"\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "Namespace": "=$flow.Namespace",
                    "DeployType": "=$flow.DeployType",
                    "ProjectID": "=$flow.ProjectID",
                    "ServiceName": "=$flow.ServiceName",
                    "InstanceID": "=$flow.InstanceID",
                    "ScriptName": "=$flow.ScriptName"
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Namespace\":\"\",\"DeployType\":\"\",\"ProjectID\":\"\",\"ServiceName\":\"\",\"InstanceID\":\"\",\"ScriptName\":\"\"}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Namespace\":\"\",\"DeployType\":\"\",\"ProjectID\":\"\",\"ServiceName\":\"\",\"InstanceID\":\"\",\"ScriptName\":\"\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"Namespace\" : \"\",\n    \"DeployType\" : \"\",\n    \"ProjectID\" : \"\",\n    \"ServiceName\" : \"\",\n    \"InstanceID\" : \"\",\n    \"ScriptName\" : \"\"\n}"
                  }
                }
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
                "MappingFields": "[{\"Name\":\"SyncDeply\",\"Type\":\"Boolean\"},{\"Name\":\"ScriptSystemEnv\",\"Type\":\"Object\"},{\"Name\":\"TargetSever\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SkipCondition": "=false",
                    "SyncDeply": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)",
                    "ScriptSystemEnv": "=f1.coalesceobject($flow.ScriptSystemEnv, f1.json2object('{}'))",
                    "TargetSever": "=f1.getsubobject($flow.ScriptSystemEnv, \"TargetServer\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeply\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"TargetSever\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"SyncDeply\":true,\"ScriptSystemEnv\":{},\"TargetSever\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeply\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"TargetSever\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"SyncDeply\":true,\"ScriptSystemEnv\":{},\"TargetSever\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"SyncDeply\",\"Type\":\"Boolean\"},{\"Name\":\"ScriptSystemEnv\",\"Type\":\"Object\"},{\"Name\":\"TargetSever\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetAgent",
            "name": "GetAgent",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"AgentID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "http://$ServiceLocatorIP$:10080/f1/locator/locate/agent/$AgentID$",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "AgentID": "=string.concat(\"agent_\", $activity[SetSystemProperties].Data.TargetSever)"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"AgentID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"AgentID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"AgentID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ExtractAgent",
            "name": "ExtractAgent",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{}},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"deployer_192.168.1.160\",\n    \"Properties\": [],\n    \"Type\": \"deployer\",\n    \"URL\": \"http://192.168.1.152:10082/f1/deployer\"\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[GetAgent].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{}},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"deployer_192.168.1.160\",\"Properties\":[],\"Type\":\"deployer\",\"URL\":\"http://192.168.1.152:10082/f1/deployer\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{}},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"ID\": \"deployer_192.168.1.160\",\n    \"Properties\": [],\n    \"Type\": \"deployer\",\n    \"URL\": \"http://192.168.1.152:10082/f1/deployer\"\n}"
                  }
                }
              }
            }
          },
          {
            "id": "LogBeforeEnvoke",
            "name": "LogBeforeEnvoke",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"@@@@@@@@@ Before Envoke @@@@@@@@@ ScriptSystemEnv : \", coerce.toString($activity[SetSystemProperties].Data.ScriptSystemEnv), \", Variables : \", coerce.toString($activity[BuildVariables].ObjectOut), \", Agent Raw : \", $activity[GetAgent].Data, \", Agent : \", coerce.toString($activity[ExtractAgent].Data))"
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
                "workingFolder": "$WorkFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"DeployType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ScriptName\",\"Type\":\"String\"},{\"Name\":\"InstanceID\",\"Type\":\"String\"}]",
                "SystemEnv": ""
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
                "Variables": "=f1.modifyobject($activity[BuildVariables].ObjectOut, \"WorkFolder\", $flow.DeployFolder)"
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
          },
          {
            "id": "BuildResources",
            "name": "BuildResources",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Path\" : \"\",\n    \"File\" : \"\"\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "Path": "/f1/deployer/resources",
                    "File": "=string.concat($flow.ProjectID, \".zip\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Path\":\"\",\"File\":\"\"}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Path\":\"\",\"File\":\"\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"File\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"Path\" : \"\",\n    \"File\" : \"\"\n}"
                  }
                }
              }
            }
          },
          {
            "id": "TableMutate",
            "name": "TableMutate",
            "description": "A simple activity for upserting/deleting data to/from table",
            "activity": {
              "ref": "#tablemutate",
              "settings": {
                "Table": {
                  "id": "809545e0-b364-11eb-9f4a-eb454cca3fe7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/TIBCOSoftware/GraphBuilder_Tools/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Deployment"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "A simple table for storing tuple"
                      },
                      "value": ""
                    },
                    {
                      "name": "schema",
                      "type": "array",
                      "required": true,
                      "display": {
                        "name": "Table Schema",
                        "description": "Schema of the table",
                        "type": "table",
                        "schema": "{\r\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\r\n    \"definitions\": {},\r\n    \"id\": \"http://example.com/example.json\",\r\n    \"items\": {\r\n        \"id\": \"/items\",\r\n        \"properties\": {\r\n            \"Name\": {\r\n                \"id\": \"/items/properties/Name\",\r\n                \"type\": \"string\"\r\n            },\r\n            \"Type\": {\r\n                \"id\": \"/items/properties/Type\",\r\n                \"type\": {\"enum\":[\"String\", \"Integer\", \"Long\", \"Double\", \"Boolean\", \"Date\", \"Object\"]}\r\n            },\r\n            \"IsKey\": {\r\n                \"id\": \"/items/properties/IsKey\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            },\r\n            \"Indexed\": {\r\n                \"id\": \"/items/properties/Indexed\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            }\r\n        },\r\n        \"type\": \"object\"\r\n    },\r\n    \"type\": \"array\"\r\n}"
                      },
                      "value": "[{\"Name\":\"AgentID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"SystemEnv\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Variables\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Resources\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/GraphBuilder_Tools/connector/simpletable/connector.json",
                  "display": {
                    "description": "This is URL file reader",
                    "category": "GraphBuilder_Tools",
                    "visible": true,
                    "smallIcon": "simpletable.png"
                  },
                  "actions": [
                    {
                      "name": "create",
                      "display": {
                        "name": "Create",
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
                  "lastUpdatedTime": 1620946644159,
                  "user": "flogo",
                  "connectorName": "Deployment",
                  "connectorDescription": " "
                },
                "Method": "upsert"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SystemEnv": "=f1.modifyobject($flow.ScriptSystemEnv, \"TargetServer\", null)",
                    "Variables": "=$activity[BuildVariables].ObjectOut",
                    "AgentID": "=$activity[ExtractAgent].Data.ID",
                    "Resources": "=$activity[BuildResources].ObjectOut"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{},\"Resources\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"New\":{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{}}}},\"Old\":{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Resources\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"New\":{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{},\"Resources\":{}},\"Old\":{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{},\"Resources\":{}}}"
                  }
                }
              }
            }
          },
          {
            "id": "ZipProjectFolder",
            "name": "ZipProjectFolder",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$WorkFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "zip -r $ProjectID$.zip ./$ProjectID$ -x ./$ProjectID$/deploy/docker/*"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "WorkFolder": "=$property[\"System.ProjectBaseFolderInt\"]"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"WorkFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"WorkFolder\":\"2\",\"ProjectID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_EnvokeDeployScriptRemote",
            "name": "Return_EnvokeDeployScriptRemote",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": true,
                  "ErrorCode": 100,
                  "Message": "Deployment information stored for agent!"
                }
              }
            }
          },
          {
            "id": "Return_EnvokeUndeployScriptRemote",
            "name": "Return_EnvokeUndeployScriptRemote",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": true,
                  "ErrorCode": 100,
                  "Message": "Deployment information stored for agent!"
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
      "id": "flow:Get_F1_Descriptor",
      "data": {
        "name": "Get F1 Descriptor",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ReadF1Descriptor",
            "to": "ResolveParameters",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ResolveParameters",
            "to": "DeserializeDscrptr",
            "type": "default"
          },
          {
            "id": 3,
            "from": "DeserializeDscrptr",
            "to": "ReturnF1Descriptor",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "ReadF1Descriptor",
            "name": "ReadF1Descriptor",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/pipeline/\" + $flow.Name + \".json\""
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
            "id": "ResolveParameters",
            "name": "ResolveParameters",
            "description": "Replace target text from input text",
            "activity": {
              "ref": "#textreplacer",
              "settings": {
                "leftToken": "$",
                "rightToken": "$",
                "replacementKeys": "[{\"Name\":\"temp.project.home\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocator\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointIP\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointPort\",\"Type\":\"String\"}]"
              },
              "input": {
                "inputDocument": "=$activity[ReadF1Descriptor].Results[0].Content",
                "replacements": {
                  "mapping": {
                    "temp.project.home": "=$flow.ProjectFolderExt + $flow.ProjectID",
                    "ID": "=$flow.ProjectID",
                    "ServiceLocator": "=$property[\"System.ServiceLocatorExternal\"]",
                    "ExternalEndpointIP": "=$property[\"System.ExternalEndpointIP\"]",
                    "ExternalEndpointPort": "10100"
                  }
                }
              },
              "schemas": {
                "input": {
                  "replacements": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"temp.project.home\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ServiceLocator\":{\"type\":\"string\"},\"ExternalEndpointIP\":{\"type\":\"string\"},\"ExternalEndpointPort\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"temp.project.home\":\"\",\"ID\":\"\",\"ServiceLocator\":\"\",\"ExternalEndpointIP\":\"\",\"ExternalEndpointPort\":\"\"}"
                  }
                },
                "settings": {
                  "replacementKeys": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"temp.project.home\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocator\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointIP\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointPort\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "DeserializeDscrptr",
            "name": "DeserializeDscrptr",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"DeploymentGpID\": {\n            \"type\": \"string\"\n        },\n        \"DataFlow\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Upstream\": {\n                        \"type\": \"string\"\n                    },\n                    \"Downstream\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        },\n        \"Components\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Runtime\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Replicas\": {\n                        \"type\": \"number\"\n                    },\n                    \"DockerImage\": {\n                        \"type\": \"string\"\n                    },\n                    \"Volumes\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"MountPoint\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Type\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"Services\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Descriptor\": {\n                        \"type\": \"string\"\n                    },\n                    \"Variables\": {\n                        \"type\": \"object\",\n                        \"properties\": {}\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Group\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"array\",\n                                    \"items\": {\n                                        \"type\": \"object\",\n                                        \"properties\": {\n                                            \"Name\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Value\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Type\": {\n                                                \"type\": \"string\"\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"System\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"DeploymentGpID\": {\n            \"type\": \"string\"\n        },\n        \"DataFlow\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Upstream\": {\n                        \"type\": \"string\"\n                    },\n                    \"Downstream\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        },\n        \"Components\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Runtime\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Replicas\": {\n                        \"type\": \"number\"\n                    },\n                    \"DockerImage\": {\n                        \"type\": \"string\"\n                    },\n                    \"Volumes\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"MountPoint\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Type\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"Services\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Descriptor\": {\n                        \"type\": \"string\"\n                    },\n                    \"Variables\": {\n                        \"type\": \"object\",\n                        \"properties\": {}\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Group\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"array\",\n                                    \"items\": {\n                                        \"type\": \"object\",\n                                        \"properties\": {\n                                            \"Name\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Value\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Type\": {\n                                                \"type\": \"string\"\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"System\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[ResolveParameters].outputDocument"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"DeploymentGpID\": {\n            \"type\": \"string\"\n        },\n        \"DataFlow\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Upstream\": {\n                        \"type\": \"string\"\n                    },\n                    \"Downstream\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        },\n        \"Components\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Runtime\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Replicas\": {\n                        \"type\": \"number\"\n                    },\n                    \"DockerImage\": {\n                        \"type\": \"string\"\n                    },\n                    \"Volumes\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"MountPoint\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Type\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"Services\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Descriptor\": {\n                        \"type\": \"string\"\n                    },\n                    \"Variables\": {\n                        \"type\": \"object\",\n                        \"properties\": {}\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Group\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"array\",\n                                    \"items\": {\n                                        \"type\": \"object\",\n                                        \"properties\": {\n                                            \"Name\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Value\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Type\": {\n                                                \"type\": \"string\"\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"System\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"DeploymentGpID\": {\n            \"type\": \"string\"\n        },\n        \"DataFlow\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Upstream\": {\n                        \"type\": \"string\"\n                    },\n                    \"Downstream\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        },\n        \"Components\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Runtime\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Replicas\": {\n                        \"type\": \"number\"\n                    },\n                    \"DockerImage\": {\n                        \"type\": \"string\"\n                    },\n                    \"Volumes\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"MountPoint\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Name\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Type\": {\n                                    \"type\": \"string\"\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"Services\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Type\": {\n                        \"type\": \"string\"\n                    },\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Descriptor\": {\n                        \"type\": \"string\"\n                    },\n                    \"Variables\": {\n                        \"type\": \"object\",\n                        \"properties\": {}\n                    },\n                    \"Properties\": {\n                        \"type\": \"array\",\n                        \"items\": {\n                            \"type\": \"object\",\n                            \"properties\": {\n                                \"Group\": {\n                                    \"type\": \"string\"\n                                },\n                                \"Value\": {\n                                    \"type\": \"array\",\n                                    \"items\": {\n                                        \"type\": \"object\",\n                                        \"properties\": {\n                                            \"Name\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Value\": {\n                                                \"type\": \"string\"\n                                            },\n                                            \"Type\": {\n                                                \"type\": \"string\"\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"System\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnF1Descriptor",
            "name": "ReturnF1Descriptor",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DeploymentGpID": "=$activity[DeserializeDscrptr].Data.DeploymentGpID",
                  "DataFlow": "=$activity[DeserializeDscrptr].Data.DataFlow",
                  "Components": "=$activity[DeserializeDscrptr].Data.Components",
                  "Services": "=$activity[DeserializeDscrptr].Data.Services",
                  "System": "=$activity[DeserializeDscrptr].Data.System"
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
              "name": "ProjectFolderExt",
              "type": "string"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "DeploymentGpID",
              "type": "string"
            },
            {
              "name": "DataFlow",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}"
              }
            },
            {
              "name": "Components",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}"
              }
            },
            {
              "name": "Services",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}"
              }
            },
            {
              "name": "System",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Deploy_F1_Descriptor_2",
      "data": {
        "name": "Deploy F1 Descriptor 2",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogF1Deploy",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "CreateDeployDescriptor",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CreateDeployDescriptor",
            "to": "EnvokeUserScript",
            "type": "default"
          },
          {
            "id": 4,
            "from": "EnvokeUserScript",
            "to": "ReturnFromUserScript",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogF1Deploy",
            "name": "LogF1Deploy",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"########## Deploy F1 Descriptor ########## ProjectID : \", $flow.ProjectID, \", SyncDeploy : \", coerce.toString($flow.SyncDeploy), \", ProjectFoldeInt : \", $flow.ProjectFolderInt, \", DeployFolder : \", $flow.DeployFolder, \", UseDefauldScript : \", coerce.toString($flow.UseDefaultScript), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv), \", F1Descriptor : \", coerce.toString($flow.F1Descriptor))"
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
                "MappingFields": "[{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "SkipCondition": "=false",
                    "ProjectFolderInt": "=$flow.ProjectFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Namespace\":\"2\",\"ProjectFolderInt\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Namespace\":\"2\",\"ProjectFolderInt\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CreateDeployDescriptor",
            "name": "CreateDeployDescriptor",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Create_Deploy_Descriptor"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "Service": "=$flow.F1Descriptor.Services[0]",
                "ProjectFolder": "=$activity[SetSystemProperties].Data.ProjectFolderInt"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Service\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Service\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "EnvokeUserScript",
            "name": "EnvokeUserScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "ScriptSystemEnv": "=f1.modifyobject($flow.ScriptSystemEnv, \"Descriptor\", $activity[CreateDeployDescriptor].Filename)",
                "ScriptName": "deploy.sh",
                "SyncDeploy": "=$flow.SyncDeploy",
                "DeployType": "=$flow.F1Descriptor.Services[0].Type",
                "ServiceName": "=$flow.F1Descriptor.Services[0].Name",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$flow.DeployFolder",
                "InstanceID": "=$flow.F1Descriptor.Services[0].Instance"
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
            "id": "ReturnFromUserScript",
            "name": "ReturnFromUserScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[EnvokeUserScript].Success",
                  "Message": "=$activity[EnvokeUserScript].Message",
                  "ErrorCode": "=$activity[EnvokeUserScript].ErrorCode"
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
                    "Message": "=string.concat(\"Deployer::Deploy F1 Descriptor - \", $error.message)"
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
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncDeploy",
              "type": "boolean"
            },
            {
              "name": "ProjectFolderInt",
              "type": "string"
            },
            {
              "name": "DeployFolder",
              "type": "string"
            },
            {
              "name": "UseDefaultScript",
              "type": "boolean"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            },
            {
              "name": "F1Descriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}",
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
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
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"F1ServiceType\",\"Type\":\"String\"}]",
                "httpHeaders": ""
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
      "id": "flow:Deploy",
      "data": {
        "name": "Deploy",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogDeploy",
            "to": "ChkUserScriptExistence",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ChkUserScriptExistence",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "GetF1Descriptor",
            "type": "expression",
            "label": "SetSystemPropertiestoStartaSubFlow",
            "value": "false==$activity[SetSystemProperties].Data.NoF1Descriptor"
          },
          {
            "id": 4,
            "from": "GetF1Descriptor",
            "to": "DeployF1Descriptor",
            "type": "default"
          },
          {
            "id": 5,
            "from": "DeployF1Descriptor",
            "to": "ReturnFromUserScript",
            "type": "default"
          },
          {
            "id": 6,
            "from": "SetSystemProperties",
            "to": "DirectEnvokeScript",
            "type": "exprOtherwise",
            "label": "SetSystemPropertiesto"
          },
          {
            "id": 7,
            "from": "DirectEnvokeScript",
            "to": "ReturnFromEnvokeScript",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogDeploy",
            "name": "LogDeploy",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"########## Deploy ########## Method : \" , $flow.Method, \" ,Action : \", $flow.Action, \" ,ProjectID : \", $flow.ProjectID, \" ,Name : \", $flow.Name, \", Instance : \", $flow.Instance, \", SyncDeploy : \", coerce.toString($flow.SyncDeploy), \", NoF1Descriptor : \", coerce.toString($flow.NoF1Descriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv), \", UserDefinedParameters : \", coerce.toString($flow.UserDefinedParameters))"
              }
            }
          },
          {
            "id": "ChkUserScriptExistence",
            "name": "ChkUserScriptExistence",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/deploy.sh\""
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
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployMethod\",\"Type\":\"String\"},{\"Name\":\"NoF1Descriptor\",\"Type\":\"Boolean\"},{\"Name\":\"UserDefinedServiceName\",\"Type\":\"String\"},{\"Name\":\"UserDefinedDescriptor\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderExt": "=$property[\"System.BaseFolderExt\"] + \"/projects/\"",
                    "ProjectFolderInt": "=string.concat($property[\"System.ProjectBaseFolderInt\"], \"/\", $flow.ProjectID)",
                    "DeployFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/\"",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "SkipCondition": "=false",
                    "SyncDeploy": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)",
                    "UseDefaultScript": "=f1.ternary(null!=$activity[ChkUserScriptExistence].Results\u0026\u00260!=array.count($activity[ChkUserScriptExistence].Results),false, true)",
                    "DeployMethod": "=f1.ternary(null!=$flow.Method, $flow.Method, \"FROM_F1_DESCRIPTOR\")",
                    "UserDefinedServiceName": "=f1.ternary(null!=$flow.UserDefinedParameters \u0026\u0026 null!=$flow.UserDefinedParameters.ServiceName, $flow.UserDefinedParameters.ServiceName, $flow.Name)",
                    "UserDefinedDescriptor": "=f1.ternary(null!=$flow.UserDefinedParameters, $flow.UserDefinedParameters.Descriptor, null)",
                    "NoF1Descriptor": "=f1.ternary(null!=$flow.NoF1Descriptor, $flow.NoF1Descriptor, false)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployMethod\":{\"type\":\"string\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"UserDefinedServiceName\":{\"type\":\"string\"},\"UserDefinedDescriptor\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"SyncDeploy\":true,\"UseDefaultScript\":true,\"DeployMethod\":\"2\",\"NoF1Descriptor\":true,\"UserDefinedServiceName\":\"2\",\"UserDefinedDescriptor\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployMethod\":{\"type\":\"string\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"UserDefinedServiceName\":{\"type\":\"string\"},\"UserDefinedDescriptor\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"SyncDeploy\":true,\"UseDefaultScript\":true,\"DeployMethod\":\"2\",\"NoF1Descriptor\":true,\"UserDefinedServiceName\":\"2\",\"UserDefinedDescriptor\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployMethod\",\"Type\":\"String\"},{\"Name\":\"NoF1Descriptor\",\"Type\":\"Boolean\"},{\"Name\":\"UserDefinedServiceName\",\"Type\":\"String\"},{\"Name\":\"UserDefinedDescriptor\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetF1Descriptor",
            "name": "GetF1Descriptor",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Get_F1_Descriptor"
              },
              "input": {
                "ProjectFolderExt": "=$activity[SetSystemProperties].Data.ProjectFolderExt",
                "ProjectID": "=$flow.ProjectID",
                "Name": "=$flow.Name"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "DeployF1Descriptor",
            "name": "DeployF1Descriptor",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Deploy_F1_Descriptor_2"
              },
              "input": {
                "F1Descriptor": {
                  "mapping": {
                    "DeploymentGpID": "=$activity[GetF1Descriptor].DeploymentGpID",
                    "DataFlow": "=$activity[GetF1Descriptor].DataFlow",
                    "Components": "=$activity[GetF1Descriptor].Components",
                    "Services": "=$activity[GetF1Descriptor].Services",
                    "System": "=$activity[GetF1Descriptor].System"
                  }
                },
                "ProjectID": "=$flow.ProjectID",
                "SyncDeploy": "=$flow.SyncDeploy",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "UseDefaultScript": "=$activity[SetSystemProperties].Data.UseDefaultScript",
                "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                "ProjectFolderInt": "=$activity[SetSystemProperties].Data.ProjectFolderInt"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}"
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
            "id": "ReturnFromUserScript",
            "name": "ReturnFromUserScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DeployF1Descriptor].Success",
                  "Message": "=$activity[DeployF1Descriptor].Message",
                  "ErrorCode": "=$activity[DeployF1Descriptor].ErrorCode"
                }
              }
            }
          },
          {
            "id": "DirectEnvokeScript",
            "name": "DirectEnvokeScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncDeploy",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                "ProjectID": "=$flow.ProjectID",
                "DeployType": "user-defined",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "ScriptName": "deploy.sh",
                "InstanceID": "=$flow.Instance",
                "ServiceName": "=$flow.Name"
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
            "id": "ReturnFromEnvokeScript",
            "name": "ReturnFromEnvokeScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DirectEnvokeScript].Success",
                  "Message": "=$activity[DirectEnvokeScript].Message",
                  "ErrorCode": "=$activity[DirectEnvokeScript].ErrorCode"
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
                    "Message": "=string.concat(\"Deployer::Deploy - \", $error.message)"
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
              "name": "NoF1Descriptor",
              "type": "boolean"
            },
            {
              "name": "Method",
              "type": "string"
            },
            {
              "name": "Action",
              "type": "string"
            },
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Instance",
              "type": "string"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            },
            {
              "name": "UserDefinedParameters",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"Method\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"UserDefinedParameters\":{\"type\":\"object\",\"properties\":{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Undeploy",
      "data": {
        "name": "Undeploy",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogUndeploy",
            "to": "ChkUserScriptExistence",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ChkUserScriptExistence",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "GetF1Descriptor",
            "type": "expression",
            "label": "SetSystemPropertiestoStartaSubFlow",
            "value": "false==$activity[SetSystemProperties].Data.NoF1Descriptor"
          },
          {
            "id": 4,
            "from": "GetF1Descriptor",
            "to": "EnvokeScriptFromF1Desc",
            "type": "default"
          },
          {
            "id": 5,
            "from": "EnvokeScriptFromF1Desc",
            "to": "ReturnFromUserScript",
            "type": "default"
          },
          {
            "id": 6,
            "from": "SetSystemProperties",
            "to": "DirectEnvokeScript",
            "type": "exprOtherwise",
            "label": "SetSystemPropertiesto"
          },
          {
            "id": 7,
            "from": "DirectEnvokeScript",
            "to": "ReturnFromEnvokeScript",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogUndeploy",
            "name": "LogUndeploy",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"########## Undeploy ########## Method : \" , $flow.Method, \" ,ProjectID\", $flow.ProjectID, \" ,SyncDeploy : \", coerce.toString($flow.SyncDeploy), \", NoF1Descriptor : \", coerce.toString($flow.NoF1Descriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv), \", UserDefinedParameters : \", coerce.toString($flow.UserDefinedParameters))"
              }
            }
          },
          {
            "id": "ChkUserScriptExistence",
            "name": "ChkUserScriptExistence",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/undeploy.sh\""
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
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployMethod\",\"Type\":\"String\"},{\"Name\":\"NoF1Descriptor\",\"Type\":\"Boolean\"},{\"Name\":\"UserDefinedServiceName\",\"Type\":\"String\"},{\"Name\":\"UserDefinedDescriptor\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderExt": "=$property[\"System.BaseFolderExt\"] + \"/projects/\"",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "DeployFolderInt": "=f1.ternary(null!=$activity[ChkUserScriptExistence].Results\u0026\u00260!=array.count($activity[ChkUserScriptExistence].Results),\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/\", \"/home/scripts\")",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "SkipCondition": "=false",
                    "SyncDeploy": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)",
                    "UseDefaultScript": "=f1.ternary(null!=$activity[ChkUserScriptExistence].Results\u0026\u00260!=array.count($activity[ChkUserScriptExistence].Results),false, true)",
                    "DeployMethod": "=f1.ternary(null!=$flow.Method, $flow.Method, \"FROM_F1_DESCRIPTOR\")",
                    "UserDefinedServiceName": "=f1.ternary(null!=$flow.UserDefinedParameters \u0026\u0026 null!=$flow.UserDefinedParameters.ServiceName, $flow.UserDefinedParameters.ServiceName, $flow.Name)",
                    "UserDefinedDescriptor": "=f1.ternary(null!=$flow.UserDefinedParameters, $flow.UserDefinedParameters.Descriptor, null)",
                    "NoF1Descriptor": "=f1.ternary(null!=$flow.NoF1Descriptor, $flow.NoF1Descriptor, false)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployMethod\":{\"type\":\"string\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"UserDefinedServiceName\":{\"type\":\"string\"},\"UserDefinedDescriptor\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"SyncDeploy\":true,\"UseDefaultScript\":true,\"DeployMethod\":\"2\",\"NoF1Descriptor\":true,\"UserDefinedServiceName\":\"2\",\"UserDefinedDescriptor\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployMethod\":{\"type\":\"string\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"UserDefinedServiceName\":{\"type\":\"string\"},\"UserDefinedDescriptor\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"SyncDeploy\":true,\"UseDefaultScript\":true,\"DeployMethod\":\"2\",\"NoF1Descriptor\":true,\"UserDefinedServiceName\":\"2\",\"UserDefinedDescriptor\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployMethod\",\"Type\":\"String\"},{\"Name\":\"NoF1Descriptor\",\"Type\":\"Boolean\"},{\"Name\":\"UserDefinedServiceName\",\"Type\":\"String\"},{\"Name\":\"UserDefinedDescriptor\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetF1Descriptor",
            "name": "GetF1Descriptor",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Get_F1_Descriptor"
              },
              "input": {
                "ProjectFolderExt": "=$activity[SetSystemProperties].Data.ProjectFolderExt",
                "ProjectID": "=$flow.ProjectID",
                "Name": "=$flow.Name"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "EnvokeScriptFromF1Desc",
            "name": "EnvokeScriptFromF1Desc",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncDeploy",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                "ProjectID": "=$flow.ProjectID",
                "DeployType": "=$activity[GetF1Descriptor].Services[0].Type",
                "ScriptName": "undeploy.sh",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "InstanceID": "=$flow.Instance",
                "ServiceName": "=$activity[SetSystemProperties].Data.UserDefinedServiceName"
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
            "id": "ReturnFromUserScript",
            "name": "ReturnFromUserScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[EnvokeScriptFromF1Desc].Message",
                  "ErrorCode": "=$activity[EnvokeScriptFromF1Desc].ErrorCode",
                  "Success": "=$activity[EnvokeScriptFromF1Desc].Success"
                }
              }
            }
          },
          {
            "id": "DirectEnvokeScript",
            "name": "DirectEnvokeScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncDeploy",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                "ProjectID": "=$flow.ProjectID",
                "DeployType": "user-defined",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "ScriptName": "undeploy.sh",
                "InstanceID": "=$flow.Instance",
                "ServiceName": "=$activity[SetSystemProperties].Data.UserDefinedServiceName"
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
            "id": "ReturnFromEnvokeScript",
            "name": "ReturnFromEnvokeScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DirectEnvokeScript].Success",
                  "Message": "=$activity[DirectEnvokeScript].Message",
                  "ErrorCode": "=$activity[DirectEnvokeScript].ErrorCode"
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
                    "Message": "=string.concat(\"Deployer::Undeploy - \", $error.message)"
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
              "name": "NoF1Descriptor",
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
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Instance",
              "type": "string"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            },
            {
              "name": "UserDefinedParameters",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"Method\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"UserDefinedParameters\":{\"type\":\"object\",\"properties\":{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Deploy_F1_Descriptor",
      "data": {
        "name": "Deploy F1 Descriptor",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogF1Deploy",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "CreateDeployDescriptor",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CreateDeployDescriptor",
            "to": "EnvokeUserScript",
            "type": "default"
          },
          {
            "id": 4,
            "from": "EnvokeUserScript",
            "to": "ReturnFromUserScript",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogF1Deploy",
            "name": "LogF1Deploy",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"########## Deploy F1 Descriptor ########## ProjectID : \", $flow.ProjectID, \", SyncDeploy : \", coerce.toString($flow.SyncDeploy), \", ProjectFoldeInt : \", $flow.ProjectFolderInt, \", DeployFolder : \", $flow.DeployFolder, \", UseDefauldScript : \", coerce.toString($flow.UseDefaultScript), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv), \", F1Descriptor : \", coerce.toString($flow.F1Descriptor))"
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
                "MappingFields": "[{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "SkipCondition": "=false",
                    "ProjectFolderInt": "=$flow.ProjectFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Namespace\":\"2\",\"ProjectFolderInt\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Namespace\":\"2\",\"ProjectFolderInt\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CreateDeployDescriptor",
            "name": "CreateDeployDescriptor",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Create_Deploy_Descriptor"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "Service": "=$flow.F1Descriptor.Services[0]",
                "ProjectFolder": "=$activity[SetSystemProperties].Data.ProjectFolderInt"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Service\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"Service\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "EnvokeUserScript",
            "name": "EnvokeUserScript",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Envoke_Script"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "ScriptSystemEnv": "=f1.modifyobject($flow.ScriptSystemEnv, \"Descriptor\", $activity[CreateDeployDescriptor].Filename)",
                "ScriptName": "deploy.sh",
                "SyncDeploy": "=$flow.SyncDeploy",
                "DeployType": "=$flow.F1Descriptor.Services[0].Type",
                "ServiceName": "=$flow.F1Descriptor.Services[0].Name",
                "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                "DeployFolder": "=$flow.DeployFolder",
                "InstanceID": "=$flow.F1Descriptor.Services[0].Instance"
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
            "id": "ReturnFromUserScript",
            "name": "ReturnFromUserScript",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[EnvokeUserScript].Success",
                  "Message": "=$activity[EnvokeUserScript].Message",
                  "ErrorCode": "=$activity[EnvokeUserScript].ErrorCode"
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
                    "Message": "=string.concat(\"Deployer::Deploy F1 Descriptor - \", $error.message)"
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
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncDeploy",
              "type": "boolean"
            },
            {
              "name": "ProjectFolderInt",
              "type": "string"
            },
            {
              "name": "DeployFolder",
              "type": "string"
            },
            {
              "name": "UseDefaultScript",
              "type": "boolean"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            },
            {
              "name": "F1Descriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Endpoint",
      "data": {
        "name": "Get Endpoint",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ReadDscrptr",
            "to": "DeserializeDscrptr",
            "type": "default"
          },
          {
            "id": 2,
            "from": "DeserializeDscrptr",
            "to": "SetSystemProperties",
            "type": "expression",
            "label": "DeserializeDscrptrtoSetSystemProperties",
            "value": "null!=$activity[DeserializeDscrptr].Data.Services\u0026\u00260!=array.count($activity[DeserializeDscrptr].Data.Services)"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "CheckEndpointLocal01",
            "type": "default"
          },
          {
            "id": 4,
            "from": "CheckEndpointLocal01",
            "to": "FetchEndpoint",
            "type": "expression",
            "label": "ReadEndpointtoFetch_Endpoint",
            "value": "0==array.count($activity[CheckEndpointLocal01].Results)"
          },
          {
            "id": 5,
            "from": "FetchEndpoint",
            "to": "CheckEndpointLocal02",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckEndpointLocal02",
            "to": "ReturnNothing02",
            "type": "expression",
            "label": "CheckEndpointLocal02to",
            "value": "0==array.count($activity[CheckEndpointLocal01].Results)"
          },
          {
            "id": 7,
            "from": "CheckEndpointLocal02",
            "to": "BuildServiceObj02",
            "type": "exprOtherwise",
            "label": "CheckEndpointLocal02toLog"
          },
          {
            "id": 8,
            "from": "BuildServiceObj02",
            "to": "ReturnServiceObj02",
            "type": "default"
          },
          {
            "id": 9,
            "from": "CheckEndpointLocal01",
            "to": "BuildServiceObj01",
            "type": "exprOtherwise",
            "label": "ReadEndpointtoCopyOfReturn_Script"
          },
          {
            "id": 10,
            "from": "BuildServiceObj01",
            "to": "ReturnServiceObj01",
            "type": "default"
          },
          {
            "id": 11,
            "from": "DeserializeDscrptr",
            "to": "ReturnNotServiceType",
            "type": "exprOtherwise",
            "label": "DeserializeDscrptr to ReturnNotServiceType"
          }
        ],
        "tasks": [
          {
            "id": "ReadDscrptr",
            "name": "ReadDscrptr",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/pipeline/descriptor.json\""
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
            "id": "DeserializeDscrptr",
            "name": "DeserializeDscrptr",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{\"Volume\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n\t\"DeploymentGpID\" : \"\",\n\t\"DataFlow\" : [{\n\t\t\t\"Upstream\" : \"\",\n\t\t\t\"Downstream\" : \"\"\n\t\t}\n\t],\n    \"Components\":[\n        {\n            \"Type\":\"\",\n            \"Runtime\":\"\",\n            \"Name\":\"\",\n\t\t\t\"Replicas\": 1,\n\t\t\t\"DockerImage\" : \"\",\n\t\t\t\"Volumes\" : [\n\t\t\t\t{\n\t\t\t\t\t\"Name\" : \"\",\n\t\t\t\t\t\"MountPoint\" : \"\"\n\t\t\t\t}\n\t\t\t],\n            \"Properties\":[\n             \t{\n                \t\"Name\":\"\",\n                \t\"Value\":\"\"\n                }\n            ]\n        }\n    ],\n   \"Services\":[\n      {\n         \"Descriptor\":\"\",\n         \"Name\":\"\",\n         \"Type\":\"\",\n         \"Variables\":{\n         },\n         \"Properties\":[\n            {\n               \"Name\":\"\",\n               \"Value\":\"\"\n            }\n        ]}\n    ],\n    \"System\" : {\n    \t\"Volume\" : [\n    \t\t{\n    \t\t\t\"Name\" : \"\",\n    \t\t\t\"Value\" : \"\"\n    \t\t}\t\n    \t]\n    }\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[ReadDscrptr].Results[0].Content"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{\"Volume\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"DeploymentGpID\":\"\",\"DataFlow\":[{\"Upstream\":\"\",\"Downstream\":\"\"}],\"Components\":[{\"Type\":\"\",\"Runtime\":\"\",\"Name\":\"\",\"Replicas\":1,\"DockerImage\":\"\",\"Volumes\":[{\"Name\":\"\",\"MountPoint\":\"\"}],\"Properties\":[{\"Name\":\"\",\"Value\":\"\"}]}],\"Services\":[{\"Descriptor\":\"\",\"Name\":\"\",\"Type\":\"\",\"Variables\":{},\"Properties\":[{\"Name\":\"\",\"Value\":\"\"}]}],\"System\":{\"Volume\":[{\"Name\":\"\",\"Value\":\"\"}]}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{\"Volume\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\n\t\"DeploymentGpID\" : \"\",\n\t\"DataFlow\" : [{\n\t\t\t\"Upstream\" : \"\",\n\t\t\t\"Downstream\" : \"\"\n\t\t}\n\t],\n    \"Components\":[\n        {\n            \"Type\":\"\",\n            \"Runtime\":\"\",\n            \"Name\":\"\",\n\t\t\t\"Replicas\": 1,\n\t\t\t\"DockerImage\" : \"\",\n\t\t\t\"Volumes\" : [\n\t\t\t\t{\n\t\t\t\t\t\"Name\" : \"\",\n\t\t\t\t\t\"MountPoint\" : \"\"\n\t\t\t\t}\n\t\t\t],\n            \"Properties\":[\n             \t{\n                \t\"Name\":\"\",\n                \t\"Value\":\"\"\n                }\n            ]\n        }\n    ],\n   \"Services\":[\n      {\n         \"Descriptor\":\"\",\n         \"Name\":\"\",\n         \"Type\":\"\",\n         \"Variables\":{\n         },\n         \"Properties\":[\n            {\n               \"Name\":\"\",\n               \"Value\":\"\"\n            }\n        ]}\n    ],\n    \"System\" : {\n    \t\"Volume\" : [\n    \t\t{\n    \t\t\t\"Name\" : \"\",\n    \t\t\t\"Value\" : \"\"\n    \t\t}\t\n    \t]\n    }\n}"
                  }
                }
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
                "MappingFields": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "DeployFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/\"",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "ServiceType": "=$activity[DeserializeDscrptr].Data.Services[0].Type"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"ServiceType\":\"2\",\"ServiceName\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"ServiceType\":\"2\",\"ServiceName\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CheckEndpointLocal01",
            "name": "CheckEndpointLocal01",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/deploy/\"+ $activity[SetSystemProperties].Data.ServiceType + \"/endpoint.json\""
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
            "id": "FetchEndpoint",
            "name": "FetchEndpoint",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$WorkFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./endpoint.sh",
                      "Execution_1": "./endpoint.sh $Namespace$ $ServiceType$ $ServiceName$"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "Namespace": "=$activity[SetSystemProperties].Data.Namespace",
                    "WorkFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                    "ServiceType": "=$activity[SetSystemProperties].Data.ServiceType",
                    "ServiceName": "=$activity[SetSystemProperties].Data.ServiceName"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"WorkFolder\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"WorkFolder\":\"2\",\"Namespace\":\"2\",\"ServiceType\":\"2\",\"ServiceName\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"WorkFolder\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceName\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CheckEndpointLocal02",
            "name": "CheckEndpointLocal02",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/deploy/\"+ $activity[SetSystemProperties].Data.ServiceType + \"/endpoint.json\""
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
            "id": "ReturnNothing02",
            "name": "ReturnNothing02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn"
            }
          },
          {
            "id": "BuildServiceObj02",
            "name": "BuildServiceObj02",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
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
                "JSONString": "=$activity[CheckEndpointLocal02].Results[0].Content"
              },
              "schemas": {
                "output": {
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
            "id": "ReturnServiceObj02",
            "name": "ReturnServiceObj02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": {
                    "mapping": {
                      "URL": "=$activity[BuildServiceObj02].Data.URL",
                      "Properties": "=$activity[BuildServiceObj02].Data.Properties",
                      "ID": "=$flow.ProjectID",
                      "Type": "=$activity[SetSystemProperties].Data.ServiceType"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "BuildServiceObj01",
            "name": "BuildServiceObj01",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
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
                "JSONString": "=$activity[CheckEndpointLocal01].Results[0].Content"
              },
              "schemas": {
                "output": {
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
            "id": "ReturnServiceObj01",
            "name": "ReturnServiceObj01",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": {
                    "mapping": {
                      "URL": "=$activity[BuildServiceObj01].Data.URL",
                      "Properties": "=$activity[BuildServiceObj01].Data.Properties",
                      "Type": "=$activity[SetSystemProperties].Data.ServiceType",
                      "ID": "=$flow.ProjectID"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "ReturnNotServiceType",
            "name": "ReturnNotServiceType",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn"
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
                "ref": "#actreturn"
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Response",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}"
          }
        },
        "explicitReply": true
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
                    "ID": "=\"deployer_\" + $property[\"System.ExternalEndpointIP\"]",
                    "Type": "deployer",
                    "URL": "=\"http://\" + $property[\"System.ServiceLocatorIP\"] + \":10082/f1/deployer\""
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
                "timeout": "500",
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
                    "ServiceType": "deployer"
                  }
                }
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
      "id": "flow:AgentServer",
      "data": {
        "name": "AgentServer",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogFindDeployment",
            "to": "FetchDeployment",
            "type": "default"
          },
          {
            "id": 2,
            "from": "FetchDeployment",
            "to": "Return",
            "type": "expression",
            "label": "TableMutatetoReturn",
            "value": "$activity[FetchDeployment].Exists"
          },
          {
            "id": 3,
            "from": "FetchDeployment",
            "to": "ReturnEmpty",
            "type": "exprOtherwise",
            "label": "FetchDeployment to ReturnEmpty"
          }
        ],
        "tasks": [
          {
            "id": "LogFindDeployment",
            "name": "LogFindDeployment",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"###### Find Deployment ###### AgentID : \" , $flow.AgentID, \", NetworkAddress : \", $flow.NetworkAddress, \", DeployedComponents : \", coerce.toString($flow.DeployedComponents))"
              }
            }
          },
          {
            "id": "FetchDeployment",
            "name": "FetchDeployment",
            "description": "A simple activity for upserting/deleting data to/from table",
            "activity": {
              "ref": "#tablemutate",
              "settings": {
                "Table": {
                  "id": "809545e0-b364-11eb-9f4a-eb454cca3fe7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/TIBCOSoftware/GraphBuilder_Tools/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Deployment"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "A simple table for storing tuple"
                      },
                      "value": ""
                    },
                    {
                      "name": "schema",
                      "type": "array",
                      "required": true,
                      "display": {
                        "name": "Table Schema",
                        "description": "Schema of the table",
                        "type": "table",
                        "schema": "{\r\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\r\n    \"definitions\": {},\r\n    \"id\": \"http://example.com/example.json\",\r\n    \"items\": {\r\n        \"id\": \"/items\",\r\n        \"properties\": {\r\n            \"Name\": {\r\n                \"id\": \"/items/properties/Name\",\r\n                \"type\": \"string\"\r\n            },\r\n            \"Type\": {\r\n                \"id\": \"/items/properties/Type\",\r\n                \"type\": {\"enum\":[\"String\", \"Integer\", \"Long\", \"Double\", \"Boolean\", \"Date\", \"Object\"]}\r\n            },\r\n            \"IsKey\": {\r\n                \"id\": \"/items/properties/IsKey\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            },\r\n            \"Indexed\": {\r\n                \"id\": \"/items/properties/Indexed\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            }\r\n        },\r\n        \"type\": \"object\"\r\n    },\r\n    \"type\": \"array\"\r\n}"
                      },
                      "value": "[{\"Name\":\"AgentID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"SystemEnv\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Variables\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Resources\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/GraphBuilder_Tools/connector/simpletable/connector.json",
                  "display": {
                    "description": "This is URL file reader",
                    "category": "GraphBuilder_Tools",
                    "visible": true,
                    "smallIcon": "simpletable.png"
                  },
                  "actions": [
                    {
                      "name": "create",
                      "display": {
                        "name": "Create",
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
                  "lastUpdatedTime": 1620946644159,
                  "user": "flogo",
                  "connectorName": "Deployment",
                  "connectorDescription": " "
                },
                "Method": "delete"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "AgentID": "=$flow.AgentID"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"New\":{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}}}},\"Old\":{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"New\":{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{}},\"Old\":{\"AgentID\":\"2\",\"SystemEnv\":{},\"Variables\":{}}}"
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
                  "Data": "=$activity[FetchDeployment].Data.Old"
                }
              }
            }
          },
          {
            "id": "ReturnEmpty",
            "name": "ReturnEmpty",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Data": "=f1.json2object(\"{}\")"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "AgentID",
              "type": "string"
            },
            {
              "name": "NetworkAddress",
              "type": "string"
            },
            {
              "name": "DeployedComponents",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Data",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"NetworkAddress\":{\"type\":\"string\"},\"DeployedComponents\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Data\":{\"type\":\"object\",\"properties\":{\"AgentID\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{\"Namespace\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"DeployType\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ServiceName\":{\"type\":\"string\"},\"InstanceID\":{\"type\":\"string\"},\"ScriptName\":{\"type\":\"string\"}}},\"SystemEnv\":{\"type\":\"object\",\"properties\":{}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Resources",
      "data": {
        "name": "Resources",
        "description": "",
        "links": [],
        "tasks": [],
        "metadata": {
          "input": [],
          "output": [],
          "fe_metadata": {}
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
