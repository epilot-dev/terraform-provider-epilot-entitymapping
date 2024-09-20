terraform {
  required_providers {
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "0.6.0"
    }
  }
}

provider "epilot-entitymapping" {
  epilot_auth = "eyJraWQiOiJ2ZFR0MGQrK1RMc2FQZ2tsQ3AzMDVGbEMxc1lOUCtUOXpsaElzMkJ3WERrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIxNzEyMTkwMy1kM2JlLTRhZTktODZiZS04YjhkZDRmYzY0ZTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfaGh6MnVJQ2xIIiwicGhvbmVfbnVtYmVyX3ZlcmlmaWVkIjp0cnVlLCJjdXN0b206aXZ5X29yZ19pZCI6IjY2IiwiY29nbml0bzp1c2VybmFtZSI6Im4uZ29lbEBlcGlsb3QuY2xvdWQiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiI4MjYwMiIsImF1ZCI6ImdqOXAwanJlaWh0cTAwY3JpNmEwZmUzMDYiLCJldmVudF9pZCI6IjlkYzhlZGM5LTk5MmMtNGY5ZS1iMzcxLTQ0NjIxZmVkYmFhMiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzI2ODY3Mjk4LCJleHAiOjE3MjY4NzA4OTgsImlhdCI6MTcyNjg2NzI5OCwiZW1haWwiOiJuLmdvZWxAZXBpbG90LmNsb3VkIn0.qLP5B4pnt7xWL7BE8Ufunm7fhnMTd16v5bWlKEjLeE9z1UGDuwrY5LLbqu7Ou8dTei3JmQuvPkEb8wNl_8leRLyXD8zjUSlTZg8dOZ2E8xyG9GxlJ0XSqgjx_pDIAC9694S6qz4QPmpXXa--QzhsVm6IwfSP1Nb-iRDd5hrp8IL1BpasOkDafYJXLyzV_V1eSEQEjVP96UBcKNtH0vaPH-Qq81ggTDOmoLFwoeIEYkyrfrT3LLBxYBAORwsfEqdBB_CRxezDUlzzWDjKqtcAeYonH25cVdWszKybc0uLTegamyt4e99dWbr4-gs028U2PETvkJqpAM03gMgI53CoLw"
}

# epilot-entitymapping_entity_mapping.entity_mapping_974a79554330440b8ce69f78636adac3:
resource "epilot-entitymapping_entity_mapping" "entity_mapping_974a79554330440b8ce69f78636adac3" {
  id = "053fcb36-97c9-46b6-8b6e-b3cd3f9d90be"
  source = {
    config = {
      journey_ref = {
        journey_id = "09946c1c-62e3-4973-bc3d-b8f500d4e7b8"
      }
    }
    type = "journey"
  }
  targets = [
    {
      allow_failure               = true
      id                          = "0291a022-ab8a-44d5-9f7b-0cfe7643f177"
      linkback_relation_attribute = "mapped_entities"
      linkback_relation_tags = [
        "customer",
        "Persönliche Informationen",
        "_hidden_contact_0291a022-ab8a-44d5-9f7b-0cfe7643f177",
      ]
      mapping_attributes = [
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                append = [
                  jsonencode(
                    {
                      email = {
                        _copy = "submission.steps[7]['Persönliche Informationen']['email']"
                      }
                    }
                  ),
                ]
                uniq = {
                  array_of_str = [
                    "email",
                  ]
                }
              }
            }
            origin = "system_recommendation"
            target = "email"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                append = [
                  jsonencode(
                    {
                      phone = {
                        _copy = "submission.steps[7]['Persönliche Informationen']['telephone']"
                      }
                    }
                  ),
                ]
                uniq = {
                  array_of_str = [
                    "phone",
                  ]
                }
              }
            }
            origin = "system_recommendation"
            target = "phone"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                append = [
                  jsonencode(
                    {
                      _id = {
                        _random = {
                          type = "nanoid"
                        }
                      }
                      _ref = "0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Installationsadresse/address"
                      _tags = [
                        "delivery",
                      ]
                      additional_info = {
                        _copy = "submission.steps[7]['Installationsadresse']['extention']"
                      }
                      city = {
                        _copy = "submission.steps[7]['Installationsadresse']['city']"
                      }
                      company_name = {
                        _copy = "submission.steps[7]['Installationsadresse']['companyName']"
                      }
                      coordinates = {
                        _copy = "submission.steps[7]['Installationsadresse']['coordinates']"
                      }
                      country = {
                        _copy = "submission.steps[7]['Installationsadresse']['countryCode']"
                      }
                      first_name = {
                        _copy = "submission.steps[7]['Installationsadresse']['firstName']"
                      }
                      last_name = {
                        _copy = "submission.steps[7]['Installationsadresse']['lastName']"
                      }
                      plot_area = {
                        _copy = "submission.steps[7]['Installationsadresse']['plotArea']"
                      }
                      plot_of_land = {
                        _copy = "submission.steps[7]['Installationsadresse']['plotOfLand']"
                      }
                      postal_code = {
                        _copy = "submission.steps[7]['Installationsadresse']['zipCode']"
                      }
                      salutation = {
                        _copy = "submission.steps[7]['Installationsadresse']['salutation']"
                      }
                      street = {
                        _copy = "submission.steps[7]['Installationsadresse']['streetName']"
                      }
                      street_number = {
                        _copy = "submission.steps[7]['Installationsadresse']['houseNumber']"
                      }
                      suburb = {
                        _copy = "submission.steps[7]['Installationsadresse']['suburb']"
                      }
                      title = {
                        _copy = "submission.steps[7]['Installationsadresse']['title']"
                      }
                    }
                  ),
                  jsonencode(
                    {
                      _id = {
                        _random = {
                          type = "nanoid"
                        }
                      }
                      _ref = "0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Abweichende Rechnungsadresse/address"
                      _tags = [
                        "billing",
                      ]
                      additional_info = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['extention']"
                      }
                      city = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['city']"
                      }
                      company_name = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['companyName']"
                      }
                      coordinates = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['coordinates']"
                      }
                      country = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['countryCode']"
                      }
                      first_name = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['firstName']"
                      }
                      last_name = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['lastName']"
                      }
                      plot_area = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['plotArea']"
                      }
                      plot_of_land = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['plotOfLand']"
                      }
                      postal_code = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['zipCode']"
                      }
                      salutation = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['salutation']"
                      }
                      street = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['streetName']"
                      }
                      street_number = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['houseNumber']"
                      }
                      suburb = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['suburb']"
                      }
                      title = {
                        _copy = "submission.steps[7]['Abweichende Rechnungsadresse']['title']"
                      }
                    }
                  ),
                ]
                uniq = {
                  array_of_str = [
                    "country",
                    "city",
                    "postal_code",
                    "street",
                    "street_number",
                    "additional_info",
                  ]
                }
              }
            }
            origin = "system_recommendation"
            target = "address"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[7]['Persönliche Informationen']['firstName']"
              }
            }
            origin = "system_recommendation"
            target = "first_name"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[7]['Persönliche Informationen']['lastName']"
              }
            }
            origin = "system_recommendation"
            target = "last_name"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[7]['Persönliche Informationen']['salutation']"
              }
            }
            origin = "system_recommendation"
            target = "salutation"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[7]['Persönliche Informationen']['title']"
              }
            }
            origin = "system_recommendation"
            target = "title"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[7]['Persönliche Informationen']['birthDate']"
              }
            }
            origin = "system_recommendation"
            target = "birthdate"
          }
        },
      ]
      name = "Contact created from block 'Persönliche Informationen' on step 'Rechnungsanschrift'"
      relation_attributes = [
        {
          mode = "append"
          source_filter = {
            limit        = 1
            relation_tag = "customer"
            schema       = "account"
            self         = false
          }
          target                     = "account"
          target_tags                = []
          target_tags_include_source = false
        },
      ]
      target_schema = "contact"
      target_unique = [
        "email.0.email",
      ]
    },
    {
      id                          = "cecd88f9-830d-4cce-8fe5-76d70d256a8d"
      linkback_relation_attribute = "mapped_entities"
      linkback_relation_tags = [
        "_hidden_opportunity_cecd88f9-830d-4cce-8fe5-76d70d256a8d",
      ]
      mapping_attributes = [
        {
          mapping_attribute_v2 = {
            operation = {
              any = jsonencode(
                {
                  _copy = [
                    "submission.journey_context.opportunity_id",
                    "opportunity._id",
                  ]
                }
              )
            }
            origin = "system_recommendation"
            target = "_id"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                set = "\"Create New Journey\""
              }
            }
            origin = "system_recommendation"
            target = "opportunity_title"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.source"
              }
            }
            origin = "system_recommendation"
            target = "source"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              any = jsonencode(
                {
                  _append = {
                    _copy = [
                      "order._tags",
                    ]
                  }
                  _uniq = true
                }
              )
            }
            origin = "system_recommendation"
            target = "_tags"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                additional_properties = jsonencode(
                  {
                    "$relation" = {
                      _append = {
                        _copy = [
                          "order.products.$relation",
                        ]
                      }
                      _uniq = [
                        "entity_id",
                      ]
                    }
                  }
                )
              }
            }
            origin = "system_recommendation"
            target = "products"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                additional_properties = jsonencode(
                  {
                    "$relation" = {
                      _append = {
                        _copy = [
                          "order.prices.$relation",
                        ]
                      }
                      _uniq = [
                        "entity_id",
                      ]
                    }
                  }
                )
              }
            }
            origin = "system_recommendation"
            target = "prices"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                additional_properties = jsonencode(
                  {
                    "$relation_ref" = {
                      _append = [
                        {
                          _id = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Abweichende Rechnungsadresse/address.refValue._id"
                          }
                          _tags = []
                          entity_id = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Abweichende Rechnungsadresse/address.entity_id"
                          }
                          path = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Abweichende Rechnungsadresse/address.attribute"
                          }
                        },
                      ]
                      _uniq = [
                        "entity_id",
                        "path",
                        "_id",
                      ]
                    }
                  }
                )
              }
            }
            origin = "system_recommendation"
            target = "billing_address"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                additional_properties = jsonencode(
                  {
                    "$relation_ref" = {
                      _append = [
                        {
                          _id = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Installationsadresse/address.refValue._id"
                          }
                          _tags = []
                          entity_id = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Installationsadresse/address.entity_id"
                          }
                          path = {
                            _copy = "refs.0291a022-ab8a-44d5-9f7b-0cfe7643f177/Rechnungsanschrift/Installationsadresse/address.attribute"
                          }
                        },
                      ]
                      _uniq = [
                        "entity_id",
                        "path",
                        "_id",
                      ]
                    }
                  }
                )
              }
            }
            origin = "system_recommendation"
            target = "delivery_address"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "consents"
              }
            }
            origin = "system_recommendation"
            target = "consents"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                append = [
                  jsonencode(
                    {
                      entity_id = {
                        _copy = "submission.journey_context.relation_id"
                      }
                    }
                  ),
                ]
                uniq = {
                  array_of_str = [
                    "entity_id",
                  ]
                }
              }
            }
            origin = "system_recommendation"
            target = "customer.$relation"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                additional_properties = jsonencode(
                  {
                    " Personen " = {
                      _copy = "submission.steps[2][' Personen ']"
                    }
                    "Add on" = {
                      _copy = "submission.steps[7]['Add on']"
                    }
                    Dachposition = {
                      _copy = "submission.steps[4]['Dachposition']"
                    }
                    Dachposition1 = {
                      _copy = "submission.steps[5]['Dachposition1']"
                    }
                    Dachschräge = {
                      _copy = "submission.steps[5]['Dachschräge']"
                    }
                    "Energy Consumption" = {
                      _copy = "submission.steps[2]['Energy Consumption']"
                    }
                    "Gefährliche Materialien" = {
                      _copy = "submission.steps[5]['Gefährliche Materialien']"
                    }
                    "Haus Typ" = {
                      _copy = "submission.steps[3]['Haus Typ']"
                    }
                    "Manuelle Eingabe" = {
                      _copy = "submission.steps[2]['Manuelle Eingabe']"
                    }
                    Postleitzahl = {
                      _copy = "submission.steps[0]['Postleitzahl']"
                    }
                    "Rechnungsadresse identisch zur Installationsadresse" = {
                      _copy = "submission.steps[7]['Rechnungsadresse identisch zur Installationsadresse']"
                    }
                    "roof size" = {
                      _copy = "submission.steps[4]['roof size']"
                    }
                  }
                )
              }
            }
            origin = "system_recommendation"
            target = "journey_data"
          }
        },
        {
          mapping_attribute_v2 = {
            operation = {
              operation_object_node = {
                copy = "submission.steps[4]['roof size']['numberInput']"
              }
            }
            origin = "user_manually"
            target = "roof_size"
          }
        },
      ]
      name = "Opportunity from Journey"
      relation_attributes = [
        {
          mode = "append"
          source_filter = {
            schema = "order"
            self   = false
          }
          target                     = "items"
          target_tags                = []
          target_tags_include_source = true
        },
        {
          mode = "append"
          source_filter = {
            limit        = 1
            relation_tag = "customer"
            schema       = "contact"
            self         = false
          }
          target = "customer"
          target_tags = [
            "customer",
          ]
          target_tags_include_source = false
        },
        {
          mode = "append"
          source_filter = {
            schema = "file"
            self   = false
          }
          target                     = "_files"
          target_tags                = []
          target_tags_include_source = true
        },
      ]
      target_schema = "opportunity"
      target_unique = [
        "_id",
      ]
    },
  ]
  version = 2
}
