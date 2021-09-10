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
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tableupsert",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tablemutate",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tablequery",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/project-flogo/legacybridge",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/contrib/function/array",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/trigger/httpredirect",
    "github.com/project-flogo/contrib/function/string",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer"
  ],
  "triggers": [
    {
      "id": "ReceiveHTTPMessage",
      "ref": "#rest",
      "settings": {
        "certFile": "",
        "enableTLS": false,
        "keyFile": "",
        "port": 10080
      },
      "handlers": [
        {
          "name": "RegisterService",
          "settings": {
            "method": "POST",
            "path": "/f1/locator/register/:service"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:RegisterService"
              },
              "input": {
                "ID": "=$.content.ID",
                "Properties": "=$.content.Properties",
                "Type": "=$.pathParams.service",
                "URL": "=$.content.URL"
              },
              "output": {
                "data": {
                  "mapping": {
                    "Response": "=$.Response"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "UnregisterService",
          "settings": {
            "method": "POST",
            "path": "/f1/locator/unregister/:service"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:UnregisterService"
              },
              "input": {
                "ID": "=$.content.ID",
                "Type": "=$.pathParams.service"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "Response": {
                      "ID": "=$.Response.ID",
                      "Type": "=$.Response.Type"
                    }
                  }
                }
              }
            }
          ]
        },
        {
          "name": "ListServices",
          "settings": {
            "method": "GET",
            "path": "/f1/locator/list/:service"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:ListServices"
              },
              "input": {
                "Type": "=$.pathParams.service"
              },
              "output": {
                "data": {
                  "mapping": {
                    "Response": "=$.Response"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "LocateService",
          "settings": {
            "method": "GET",
            "path": "/f1/locator/locate/:ServiceType/:serviceID"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:LocateService"
              },
              "input": {
                "ServiceID": "=$.pathParams.serviceID",
                "ServiceType": "=$.pathParams.ServiceType"
              },
              "output": {
                "data": "=$.Response"
              }
            }
          ]
        },
        {
          "name": "GetMetadata",
          "settings": {
            "method": "GET",
            "path": "/f1/metadata/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:GetMetadata"
              },
              "input": {
                "Name": "=$.pathParams.name"
              },
              "output": {
                "data": {
                  "mapping": {
                    "Name": "=$.Name",
                    "Value": "=$.Value"
                  }
                }
              }
            }
          ]
        }
      ]
    },
    {
      "id": "HTTPRedirect",
      "ref": "#httpredirect",
      "settings": {
        "Port": "10081"
      },
      "handlers": [
        {
          "name": "HTTPRedirect",
          "settings": {
            "Path": "/"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:HTTPRedirect"
              },
              "input": {
                "RequestURL": "=$.RequestURL"
              },
              "output": {
                "data": {
                  "mapping": {
                    "RedirectURL": "=$.RedirectURL"
                  }
                }
              }
            }
          ],
          "schemas": {
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"RedirectURL\" : \"\"\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RedirectURL\":{\"type\":\"string\"}}}"
              }
            }
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:RegisterService",
      "data": {
        "name": "RegisterService",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogRegisterService",
            "to": "TableUpsert",
            "type": "default"
          },
          {
            "id": 2,
            "from": "TableUpsert",
            "to": "Log",
            "type": "default"
          },
          {
            "id": 3,
            "from": "Log",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogRegisterService",
            "name": "LogRegisterService",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Register Service ^^^^^^ ID : \", $flow.ID, \", Type : \", $flow.Type, \"URL : \", $flow.URL, \", Properties : \", coerce.toString($flow.Properties))"
              }
            }
          },
          {
            "id": "TableUpsert",
            "name": "TableUpsert",
            "description": "A simple activity for upserting data to table",
            "activity": {
              "ref": "#tableupsert",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                }
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ServiceID": "=$flow.ID",
                    "ServiceType": "=$flow.Type",
                    "URL": "=$flow.URL",
                    "Properties": "=$flow.Properties"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "Log",
            "name": "Log",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString($activity[TableUpsert].Data)",
                "addDetails": false
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
                  "Response": {
                    "mapping": {
                      "ID": "=$activity[TableUpsert].Data.ServiceID",
                      "Type": "=$activity[TableUpsert].Data.ServiceType",
                      "URL": "=$activity[TableUpsert].Data.URL",
                      "Properties": "=$activity[TableUpsert].Data.Properties"
                    }
                  }
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Type",
              "type": "string"
            },
            {
              "name": "URL",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:UnregisterService",
      "data": {
        "name": "UnregisterService",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogUnregisterService",
            "to": "TableMutate",
            "type": "default"
          },
          {
            "id": 2,
            "from": "TableMutate",
            "to": "Return",
            "type": "expression",
            "label": "TableMutatetoReturn",
            "value": "$activity[TableMutate].Exists"
          },
          {
            "id": 3,
            "from": "TableMutate",
            "to": "NotFound",
            "type": "exprOtherwise",
            "label": "TableMutatetoReturn1"
          }
        ],
        "tasks": [
          {
            "id": "LogUnregisterService",
            "name": "LogUnregisterService",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Unregister Service ^^^^^^ ID : \", $flow.ID, \", Type : \", $flow.Type)"
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
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Method": "delete"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ServiceID": "=$flow.ID",
                    "ServiceType": "=$flow.Type"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
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
                  "Response": {
                    "mapping": {
                      "ID": "=$activity[TableMutate].Data.ServiceID",
                      "Type": "=$activity[TableMutate].Data.ServiceType"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "NotFound",
            "name": "NotFound",
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
            },
            {
              "name": "Type",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Response",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:ListServices",
      "data": {
        "name": "ListServices",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogListServices",
            "to": "TableQuery",
            "type": "default"
          },
          {
            "id": 2,
            "from": "TableQuery",
            "to": "LogQueryResult",
            "type": "default"
          },
          {
            "id": 3,
            "from": "LogQueryResult",
            "to": "Mapping",
            "type": "expression",
            "label": "LogQueryResulttoMapping",
            "value": "0!=array.count($activity[TableQuery].Data)"
          },
          {
            "id": 4,
            "from": "Mapping",
            "to": "Return",
            "type": "default"
          },
          {
            "id": 5,
            "from": "LogQueryResult",
            "to": "Return_Nothing",
            "type": "exprOtherwise",
            "label": "LogQueryResult to Return_Nothing"
          }
        ],
        "tasks": [
          {
            "id": "LogListServices",
            "name": "LogListServices",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ List Services ^^^^^^  Type : \", $flow.Type)"
              }
            }
          },
          {
            "id": "TableQuery",
            "name": "TableQuery",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Indices": "ServiceType"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ServiceType": "=$flow.Type"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}",
                    "fe_metadata": "[{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}]"
                  }
                }
              }
            }
          },
          {
            "id": "LogQueryResult",
            "name": "LogQueryResult",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Query Result ^^^^^^ TableQuery.Data : \", coerce.toString($activity[TableQuery].Data))"
              }
            }
          },
          {
            "id": "Mapping",
            "name": "Mapping",
            "description": "Mapping field from input to output",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[TableQuery].Data",
              "accumulate": false
            },
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": true,
                "MappingFields": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ArraySize": "=array.count($activity[TableQuery].Data)",
                    "SkipCondition": "=false",
                    "Type": "=$iteration[value].ServiceType",
                    "ID": "=$iteration[value].ServiceID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"},\"ArraySize\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"Type\":\"2\",\"SkipCondition\":false,\"ArraySize\":2}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"ID\":\"2\",\"Type\":\"2\"}]"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"}]"
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
                  "Response": "=$activity[Mapping].Data"
                }
              }
            }
          },
          {
            "id": "Return_Nothing",
            "name": "Return_Nothing",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": {
                    "mapping": []
                  }
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Type",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Response",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:LocateService",
      "data": {
        "name": "LocateService",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogLocateService",
            "to": "FindServiceInLocator",
            "type": "default"
          },
          {
            "id": 2,
            "from": "FindServiceInLocator",
            "to": "LogServiceFound",
            "type": "expression",
            "label": "TableQuerytoReturn",
            "value": "$activity[FindServiceInLocator].Exists"
          },
          {
            "id": 3,
            "from": "LogServiceFound",
            "to": "Return",
            "type": "default"
          },
          {
            "id": 4,
            "from": "FindServiceInLocator",
            "to": "LogServiceNotFound",
            "type": "exprOtherwise",
            "label": "TableQuerytoReturn1"
          },
          {
            "id": 5,
            "from": "LogServiceNotFound",
            "to": "LogDeligateToDeployer",
            "type": "expression",
            "label": "LogServiceNotFoundtoLogDeligateToDeployer",
            "value": "\"agent\"!=$flow.ServiceType\u0026\u0026\"deployer\"!=$flow.ServiceType\u0026\u0026\"builder\"!=$flow.ServiceType\u0026\u0026\"locator\"!=$flow.ServiceType\u0026\u0026\"projectmgr\"!=$flow.ServiceType"
          },
          {
            "id": 6,
            "from": "LogDeligateToDeployer",
            "to": "FindDeployer",
            "type": "default"
          },
          {
            "id": 7,
            "from": "FindDeployer",
            "to": "FindEndpoint",
            "type": "default"
          },
          {
            "id": 8,
            "from": "FindEndpoint",
            "to": "ResponseFromDeployer",
            "type": "default"
          },
          {
            "id": 9,
            "from": "ResponseFromDeployer",
            "to": "SaveServiceInfo",
            "type": "expression",
            "label": "JSONDeserializertoReturnFromDeployer",
            "value": "$activity[FindEndpoint].Success\u0026\u0026null!=$activity[ResponseFromDeployer].Data.ID"
          },
          {
            "id": 10,
            "from": "SaveServiceInfo",
            "to": "ReturnFromDeployer",
            "type": "default"
          },
          {
            "id": 11,
            "from": "ResponseFromDeployer",
            "to": "NotFoundFromDeployer",
            "type": "exprOtherwise",
            "label": "ResponseFromDeployer to NotFoundFromDeployer"
          },
          {
            "id": 12,
            "from": "LogServiceNotFound",
            "to": "ReturnNothing",
            "type": "exprOtherwise",
            "label": "LogServiceNotFound to ReturnNothing"
          }
        ],
        "tasks": [
          {
            "id": "LogLocateService",
            "name": "LogLocateService",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ LocateService ^^^^^^ ServiceType : \", $flow.ServiceType, \", ServiceID : \", $flow.ServiceID)"
              }
            }
          },
          {
            "id": "FindServiceInLocator",
            "name": "FindServiceInLocator",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Indices": "ServiceID ServiceType"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ServiceID": "=$flow.ServiceID",
                    "ServiceType": "=$flow.ServiceType"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}",
                    "fe_metadata": "[{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}]"
                  }
                }
              }
            }
          },
          {
            "id": "LogServiceFound",
            "name": "LogServiceFound",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ ServiceFound ^^^^^^ ServiceID : \", $flow.ServiceID)"
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
                  "Response": {
                    "mapping": {
                      "ID": "=$activity[FindServiceInLocator].Data[0].ServiceID",
                      "Type": "=$activity[FindServiceInLocator].Data[0].ServiceType",
                      "URL": "=$activity[FindServiceInLocator].Data[0].URL",
                      "Properties": "=$activity[FindServiceInLocator].Data[0].Properties"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "LogServiceNotFound",
            "name": "LogServiceNotFound",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Service Not Found ^^^^^^ ServiceID : \", $flow.ServiceID)"
              }
            }
          },
          {
            "id": "LogDeligateToDeployer",
            "name": "LogDeligateToDeployer",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ DeligateToDeployer ^^^^^^ ServiceType : \", $flow.ServiceType)"
              }
            }
          },
          {
            "id": "FindDeployer",
            "name": "FindDeployer",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Indices": "ServiceType"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ServiceType": "deployer"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}",
                    "fe_metadata": "[{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}]"
                  }
                }
              }
            }
          },
          {
            "id": "FindEndpoint",
            "name": "FindEndpoint",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "http://deployer:10082/f1/deployer/endpoint/$ServiceID$",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceID": "=$flow.ServiceID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ServiceID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ResponseFromDeployer",
            "name": "ResponseFromDeployer",
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
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[FindEndpoint].Data"
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
            "id": "SaveServiceInfo",
            "name": "SaveServiceInfo",
            "description": "A simple activity for upserting data to table",
            "activity": {
              "ref": "#tableupsert",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                }
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ServiceID": "=$activity[ResponseFromDeployer].Data.ID",
                    "ServiceType": "=$activity[ResponseFromDeployer].Data.Type",
                    "URL": "=$activity[ResponseFromDeployer].Data.URL",
                    "Properties": "=$activity[ResponseFromDeployer].Data.Properties"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "ReturnFromDeployer",
            "name": "ReturnFromDeployer",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": {
                    "mapping": {
                      "ID": "=$activity[SaveServiceInfo].Data.ServiceID",
                      "Type": "=$activity[SaveServiceInfo].Data.ServiceType",
                      "URL": "=$activity[SaveServiceInfo].Data.URL",
                      "Properties": "=$activity[SaveServiceInfo].Data.Properties"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "NotFoundFromDeployer",
            "name": "NotFoundFromDeployer",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": "=$activity[ResponseFromDeployer].Data"
                }
              }
            }
          },
          {
            "id": "ReturnNothing",
            "name": "ReturnNothing",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Response": "=f1.json2object(\"{}\")"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "ServiceID",
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
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"},\"ServiceID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:HTTPRedirect",
      "data": {
        "name": "HTTPRedirect",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "GetServiceType",
            "to": "Log",
            "type": "default"
          },
          {
            "id": 2,
            "from": "Log",
            "to": "TableQuery",
            "type": "expression",
            "label": "LogtoMapping",
            "value": "\"ingerence\"!=$activity[GetServiceType].Data.ServiceType"
          },
          {
            "id": 3,
            "from": "TableQuery",
            "to": "PrintURL01",
            "type": "default"
          },
          {
            "id": 4,
            "from": "PrintURL01",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 5,
            "from": "Log",
            "to": "AccessByID",
            "type": "expression",
            "label": "LogtoAccessByID",
            "value": "\"ingerence\"!=$activity[GetServiceType].Data.ServiceType"
          },
          {
            "id": 6,
            "from": "AccessByID",
            "to": "QueryByTypeNdID",
            "type": "default"
          },
          {
            "id": 7,
            "from": "QueryByTypeNdID",
            "to": "PrintURL02",
            "type": "default"
          },
          {
            "id": 8,
            "from": "PrintURL02",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "GetServiceType",
            "name": "GetServiceType",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ServiceType": "=array.get(string.split($flow.RequestURL, \"/\"), 2)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Log",
            "name": "Log",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$flow.RequestURL",
                "addDetails": false
              }
            }
          },
          {
            "id": "TableQuery",
            "name": "TableQuery",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Indices": "ServiceType"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ServiceType": "=$activity[GetServiceType].Data.ServiceType"
                  }
                }
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}",
                    "fe_metadata": "[{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}]"
                  }
                }
              }
            }
          },
          {
            "id": "PrintURL01",
            "name": "PrintURL01",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$activity[TableQuery].Data[0].URL",
                "addDetails": false
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
                  "RedirectURL": "=$activity[TableQuery].Data[0].URL+string.substringAfter($flow.RequestURL, $activity[GetServiceType].Data.ServiceType)"
                }
              }
            }
          },
          {
            "id": "AccessByID",
            "name": "AccessByID",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ServiceType": "=array.get(string.split($flow.RequestURL, \"/\"), 2)",
                    "ServiceID": "=array.get(string.split($flow.RequestURL, \"/\"), 3)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"},\"ServiceID\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\",\"ServiceID\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceType\":{\"type\":\"string\"},\"ServiceID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceType\":\"2\",\"ServiceID\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceType\",\"Type\":\"String\"},{\"Name\":\"ServiceID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "QueryByTypeNdID",
            "name": "QueryByTypeNdID",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c6215040-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Services"
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
                      "value": "[{\"Name\":\"ServiceID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"ServiceType\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"yes\"},{\"Name\":\"Plateform\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"URL\",\"Type\":\"String\",\"IsKey\":\"no\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374841324,
                  "user": "flogo",
                  "connectorName": "Services",
                  "connectorDescription": " "
                },
                "Indices": "ServiceID ServiceType"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ServiceType": "=$activity[AccessByID].Data.ServiceType",
                    "ServiceID": "=$activity[AccessByID].Data.ServiceID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceID\":\"2\",\"ServiceType\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ServiceID\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Plateform\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}",
                    "fe_metadata": "[{\"ServiceID\":\"2\",\"ServiceType\":\"2\",\"Plateform\":\"2\",\"URL\":\"2\",\"Properties\":{}}]"
                  }
                }
              }
            }
          },
          {
            "id": "PrintURL02",
            "name": "PrintURL02",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=$activity[QueryByTypeNdID].Data[0].URL",
                "addDetails": false
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
                  "RedirectURL": "=$activity[QueryByTypeNdID].Data[0].URL"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "RequestURL",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "RedirectURL",
              "type": "string"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RequestURL\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RedirectURL\":{\"type\":\"string\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:GetMetadata",
      "data": {
        "name": "GetMetadata",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogGetMetadata",
            "to": "TableQuery",
            "type": "default"
          },
          {
            "id": 2,
            "from": "TableQuery",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogGetMetadata",
            "name": "LogGetMetadata",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"^^^^^^ Get Metadata ^^^^^^ Name : \", $flow.Name)"
              }
            }
          },
          {
            "id": "TableQuery",
            "name": "TableQuery",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "c62065e0-4aff-11eb-9ffb-8db6499bd1d7",
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "Metadata"
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
                      "value": "[{\"Name\":\"Name\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Value\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609374835534,
                  "user": "flogo",
                  "connectorName": "Metadata",
                  "connectorDescription": " "
                },
                "Indices": "Name"
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"Value\":{}}"
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
              "ref": "#actreturn"
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Name",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Value",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{}}}}"
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
