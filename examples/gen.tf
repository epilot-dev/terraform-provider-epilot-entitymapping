# epilot-entitymapping_entity_mapping.new_entity_mapping:
resource "epilot-entitymapping_entity_mapping" "new_entity_mapping" {
    id      = "b3479ba8-e49d-40ba-868a-88b7cda6dbcc"
    source  = {
        config = {
            journey_ref = {
                journey_id = "08d948c0-49e9-11ef-a0cf-23e85ed2a58f"
            }
        }
        type   = "journey"
    }
    targets = [
        {
            allow_failure               = true
            id                          = "23132a1d-8a73-487e-996b-2c6d4bf1a8e8"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = [
                "customer",
                "Persönliche Informationen",
            ]
            mapping_attributes          = [
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                append = [
                                    jsonencode(
                                        {
                                            email = {
                                                _copy = "submission.steps[0]['Persönliche Informationen']['email']"
                                            }
                                        }
                                    ),
                                ]
                                uniq   = {
                                    array_of_str = [
                                        "email",
                                    ]
                                }
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "email"
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
                                                _copy = "submission.steps[0]['Persönliche Informationen']['telephone']"
                                            }
                                        }
                                    ),
                                ]
                                uniq   = {
                                    array_of_str = [
                                        "phone",
                                    ]
                                }
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "phone"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                append = [
                                    jsonencode(
                                        {
                                            _id             = {
                                                _random = {
                                                    type = "nanoid"
                                                }
                                            }
                                            _ref            = "23132a1d-8a73-487e-996b-2c6d4bf1a8e8/Persönliche Details/Adresse/address"
                                            _tags           = []
                                            additional_info = {
                                                _copy = "submission.steps[0]['Adresse']['extention']"
                                            }
                                            city            = {
                                                _copy = "submission.steps[0]['Adresse']['city']"
                                            }
                                            company_name    = {
                                                _copy = "submission.steps[0]['Adresse']['companyName']"
                                            }
                                            coordinates     = {
                                                _copy = "submission.steps[0]['Adresse']['coordinates']"
                                            }
                                            country         = {
                                                _copy = "submission.steps[0]['Adresse']['countryCode']"
                                            }
                                            first_name      = {
                                                _copy = "submission.steps[0]['Adresse']['firstName']"
                                            }
                                            last_name       = {
                                                _copy = "submission.steps[0]['Adresse']['lastName']"
                                            }
                                            plot_area       = {
                                                _copy = "submission.steps[0]['Adresse']['plotArea']"
                                            }
                                            plot_of_land    = {
                                                _copy = "submission.steps[0]['Adresse']['plotOfLand']"
                                            }
                                            postal_code     = {
                                                _copy = "submission.steps[0]['Adresse']['zipCode']"
                                            }
                                            salutation      = {
                                                _copy = "submission.steps[0]['Adresse']['salutation']"
                                            }
                                            street          = {
                                                _copy = "submission.steps[0]['Adresse']['streetName']"
                                            }
                                            street_number   = {
                                                _copy = "submission.steps[0]['Adresse']['houseNumber']"
                                            }
                                            suburb          = {
                                                _copy = "submission.steps[0]['Adresse']['suburb']"
                                            }
                                            title           = {
                                                _copy = "submission.steps[0]['Adresse']['title']"
                                            }
                                        }
                                    ),
                                ]
                                uniq   = {
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
                        origin    = "system_recommendation"
                        target    = "address"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Persönliche Informationen']['firstName']"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "first_name"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Persönliche Informationen']['lastName']"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "last_name"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Persönliche Informationen']['salutation']"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "salutation"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Persönliche Informationen']['title']"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "title"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Persönliche Informationen']['birthDate']"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "birthdate"
                    }
                },
            ]
            name                        = "Kontakt erstellt aus Block 'Persönliche Informationen' auf Schritt 'Persönliche Details'"
            relation_attributes         = [
                {
                    mode                       = "append"
                    source_filter              = {
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
            target_schema               = "contact"
            target_unique               = [
                "email.0.email",
            ]
        },
        {
            id                          = "252e73dd-325e-4f5e-aa4a-57267705ae0b"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = []
            mapping_attributes          = [
                {
                    mapping_attribute_v2 = {
                        operation = {
                            any = jsonencode(
                                {
                                    _copy = [
                                        "journey_context.opportunity_id",
                                        "opportunity._id",
                                    ]
                                }
                            )
                        }
                        origin    = "system_recommendation"
                        target    = "_id"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                set = "\"Journey To Export\""
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "opportunity_title"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.source"
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "source"
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
                                    _uniq   = true
                                }
                            )
                        }
                        origin    = "system_recommendation"
                        target    = "_tags"
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
                                            _uniq   = [
                                                "entity_id",
                                            ]
                                        }
                                    }
                                )
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "products"
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
                                            _uniq   = [
                                                "entity_id",
                                            ]
                                        }
                                    }
                                )
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "prices"
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
                                                    _id       = {
                                                        _copy = "refs.23132a1d-8a73-487e-996b-2c6d4bf1a8e8/Persönliche Details/Adresse/address.refValue._id"
                                                    }
                                                    entity_id = {
                                                        _copy = "refs.23132a1d-8a73-487e-996b-2c6d4bf1a8e8/Persönliche Details/Adresse/address.entity_id"
                                                    }
                                                    path      = {
                                                        _copy = "refs.23132a1d-8a73-487e-996b-2c6d4bf1a8e8/Persönliche Details/Adresse/address.attribute"
                                                    }
                                                },
                                            ]
                                            _uniq   = [
                                                "entity_id",
                                                "path",
                                                "_id",
                                            ]
                                        }
                                    }
                                )
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "address"
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
                                uniq   = {
                                    array_of_str = [
                                        "entity_id",
                                    ]
                                }
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "customer.$relation"
                    }
                },
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                additional_properties = jsonencode(
                                    {
                                        Entitätssuche          = {
                                            _copy = "submission.steps[0]['Entitätssuche']"
                                        }
                                        "Preferred Energytype" = {
                                            _copy = "submission.steps[1]['Preferred Energytype']"
                                        }
                                        "Your Request Block"   = {
                                            _copy = "submission.steps[1]['Your Request Block']"
                                        }
                                    }
                                )
                            }
                        }
                        origin    = "system_recommendation"
                        target    = "journey_data"
                    }
                },
            ]
            name                        = "Opportunity aus Journey"
            relation_attributes         = [
                {
                    mode                       = "append"
                    source_filter              = {
                        schema = "order"
                        self   = false
                    }
                    target                     = "items"
                    target_tags                = []
                    target_tags_include_source = true
                },
                {
                    mode                       = "append"
                    source_filter              = {
                        limit        = 1
                        relation_tag = "customer"
                        schema       = "contact"
                        self         = false
                    }
                    target                     = "customer"
                    target_tags                = []
                    target_tags_include_source = false
                },
                {
                    mode                       = "append"
                    source_filter              = {
                        schema = "file"
                        self   = false
                    }
                    target                     = "_files"
                    target_tags                = []
                    target_tags_include_source = true
                },
            ]
            target_schema               = "opportunity"
            target_unique               = [
                "_id",
            ]
        },
        {
            allow_failure               = true
            id                          = "60cc81de-e44e-4510-b8df-153e13cc0a24"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = [
                "lookup_0",
                "CONTRACT CUSTOMER",
            ]
            mapping_attributes          = [
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Entitätssuche']['entity']['customer'][0]['_id']"
                            }
                        }
                        origin    = "entity_updating_system_recommendation"
                        target    = "_id"
                    }
                },
            ]
            name                        = "Kontakt erstellt aus Block 'Entitätssuche' auf Schritt 'Persönliche Details'"
            relation_attributes         = []
            target_schema               = "contact"
            target_unique               = [
                "_id",
            ]
        },
        {
            allow_failure               = true
            id                          = "a54a0115-6ddf-4e8d-a5fc-412d586a3039"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = [
                "lookup_0",
                "111Persönliche Details | Entitätssuche",
            ]
            mapping_attributes          = [
                {
                    mapping_attribute_v2 = {
                        operation = {
                            operation_object_node = {
                                copy = "submission.steps[0]['Entitätssuche']['entity']['_id']"
                            }
                        }
                        origin    = "entity_updating_system_recommendation"
                        target    = "_id"
                    }
                },
            ]
            name                        = "Entity Mapping: Persönliche Details | Entitätssuche"
            relation_attributes         = [
                {
                    mode                       = "append"
                    source_filter              = {
                        limit        = 1
                        relation_tag = "CONTRACT CUSTOMER"
                        schema       = "contact"
                        self         = false
                    }
                    target                     = "customer"
                    target_tags                = []
                    target_tags_include_source = false
                },
            ]
            target_schema               = "contract"
            target_unique               = [
                "_id",
            ]
        },
    ]
}
