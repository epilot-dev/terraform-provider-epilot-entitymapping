# epilot-entitymapping_entity_mapping.entity_mapping_974a79554330440b8ce69f78636adac3:
resource "epilot-entitymapping_entity_mapping" "entity_mapping_974a79554330440b8ce69f78636adac3" {
    id      = "305af875-3883-40d3-9272-48ae16ea500a"
    source  = {
        config = {
            journey_ref = {
                journey_id = "f6e1d960-7fd6-11ef-bcfa-efcd501e9a89"
            }
        }
        type   = "journey"
    }
    targets = [
        {
            allow_failure               = true
            id                          = "c450ea10-752e-42ff-a51c-2986bc01003b"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = [
                "customer",
                "Contact Info",
                "_hidden_contact_c450ea10-752e-42ff-a51c-2986bc01003b",
            ]
            mapping_attributes          = [
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        email = {
                                            _copy = "submission.steps[4]['Contact Info']['email']"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "email",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "email"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        phone = {
                                            _copy = "submission.steps[4]['Contact Info']['telephone']"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "phone",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "phone"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        _id             = {
                                            _random = {
                                                type = "nanoid"
                                            }
                                        }
                                        _ref            = "c450ea10-752e-42ff-a51c-2986bc01003b/Persönliche Informationen/Address/address"
                                        _tags           = [
                                            "billing",
                                            "delivery",
                                        ]
                                        additional_info = {
                                            _copy = "submission.steps[4]['Address']['extention']"
                                        }
                                        city            = {
                                            _copy = "submission.steps[4]['Address']['city']"
                                        }
                                        company_name    = {
                                            _copy = "submission.steps[4]['Address']['companyName']"
                                        }
                                        coordinates     = {
                                            _copy = "submission.steps[4]['Address']['coordinates']"
                                        }
                                        country         = {
                                            _copy = "submission.steps[4]['Address']['countryCode']"
                                        }
                                        first_name      = {
                                            _copy = "submission.steps[4]['Address']['firstName']"
                                        }
                                        last_name       = {
                                            _copy = "submission.steps[4]['Address']['lastName']"
                                        }
                                        plot_area       = {
                                            _copy = "submission.steps[4]['Address']['plotArea']"
                                        }
                                        plot_of_land    = {
                                            _copy = "submission.steps[4]['Address']['plotOfLand']"
                                        }
                                        postal_code     = {
                                            _copy = "submission.steps[4]['Address']['zipCode']"
                                        }
                                        salutation      = {
                                            _copy = "submission.steps[4]['Address']['salutation']"
                                        }
                                        street          = {
                                            _copy = "submission.steps[4]['Address']['streetName']"
                                        }
                                        street_number   = {
                                            _copy = "submission.steps[4]['Address']['houseNumber']"
                                        }
                                        suburb          = {
                                            _copy = "submission.steps[4]['Address']['suburb']"
                                        }
                                        title           = {
                                            _copy = "submission.steps[4]['Address']['title']"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "country",
                                    "city",
                                    "postal_code",
                                    "street",
                                    "street_number",
                                    "additional_info",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "address"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['firstName']"
                            }
                            origin    = "system_recommendation"
                            target    = "first_name"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['lastName']"
                            }
                            origin    = "system_recommendation"
                            target    = "last_name"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['salutation']"
                            }
                            origin    = "system_recommendation"
                            target    = "salutation"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['title']"
                            }
                            origin    = "system_recommendation"
                            target    = "title"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['birthDate']"
                            }
                            origin    = "system_recommendation"
                            target    = "birthdate"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        _id   = {
                                            _random = {
                                                type = "nanoid"
                                            }
                                        }
                                        _ref  = "Zahlung/Payment Method/payment"
                                        _tags = []
                                        data  = {
                                            _copy = "submission.steps[5]['Payment Method']['data']"
                                        }
                                        type  = {
                                            _copy = "submission.steps[5]['Payment Method']['type']"
                                        }
                                    },
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "payment"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                   epilot-taxonomy_taxonomy_classification.taxonomy_classification_cb92a48f0b1b49b3af2d73c54cadf186.id,
                                   epilot-taxonomy_taxonomy_classification.taxonomy_classification_361ea610da2e488a809d12288d8ca28a.id,
                                ]
                            }
                            origin    = "user_manually"
                            target    = "_purpose"
                        }
                    )
                },
            ]
            name                        = "Contact created from block 'Contact Info' on step 'Persönliche Informationen'"
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
            allow_failure               = true
            id                          = "eadce970-5cd1-404a-96a7-e5a8f54e52a7"
            linkback_relation_attribute = "mapped_entities"
            linkback_relation_tags      = [
                "_hidden_order_eadce970-5cd1-404a-96a7-e5a8f54e52a7",
            ]
            mapping_attributes          = [
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        _id             = {
                                            _random = {
                                                type = "nanoid"
                                            }
                                        }
                                        _ref            = "eadce970-5cd1-404a-96a7-e5a8f54e52a7/Persönliche Informationen/Address/billing_address"
                                        _tags           = [
                                            "billing",
                                            "delivery",
                                        ]
                                        additional_info = {
                                            _copy = "submission.steps[4]['Address']['extention']"
                                        }
                                        city            = {
                                            _copy = "submission.steps[4]['Address']['city']"
                                        }
                                        company_name    = {
                                            _copy = "submission.steps[4]['Address']['companyName']"
                                        }
                                        coordinates     = {
                                            _copy = "submission.steps[4]['Address']['coordinates']"
                                        }
                                        country         = {
                                            _copy = "submission.steps[4]['Address']['countryCode']"
                                        }
                                        first_name      = {
                                            _copy = "submission.steps[4]['Address']['firstName']"
                                        }
                                        last_name       = {
                                            _copy = "submission.steps[4]['Address']['lastName']"
                                        }
                                        plot_area       = {
                                            _copy = "submission.steps[4]['Address']['plotArea']"
                                        }
                                        plot_of_land    = {
                                            _copy = "submission.steps[4]['Address']['plotOfLand']"
                                        }
                                        postal_code     = {
                                            _copy = "submission.steps[4]['Address']['zipCode']"
                                        }
                                        salutation      = {
                                            _copy = "submission.steps[4]['Address']['salutation']"
                                        }
                                        street          = {
                                            _copy = "submission.steps[4]['Address']['streetName']"
                                        }
                                        street_number   = {
                                            _copy = "submission.steps[4]['Address']['houseNumber']"
                                        }
                                        suburb          = {
                                            _copy = "submission.steps[4]['Address']['suburb']"
                                        }
                                        title           = {
                                            _copy = "submission.steps[4]['Address']['title']"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "country",
                                    "city",
                                    "postal_code",
                                    "street",
                                    "street_number",
                                    "additional_info",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "billing_address"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        _id             = {
                                            _random = {
                                                type = "nanoid"
                                            }
                                        }
                                        _ref            = "eadce970-5cd1-404a-96a7-e5a8f54e52a7/Persönliche Informationen/Address/delivery_address"
                                        _tags           = [
                                            "billing",
                                            "delivery",
                                        ]
                                        additional_info = {
                                            _copy = "submission.steps[4]['Address']['extention']"
                                        }
                                        city            = {
                                            _copy = "submission.steps[4]['Address']['city']"
                                        }
                                        company_name    = {
                                            _copy = "submission.steps[4]['Address']['companyName']"
                                        }
                                        coordinates     = {
                                            _copy = "submission.steps[4]['Address']['coordinates']"
                                        }
                                        country         = {
                                            _copy = "submission.steps[4]['Address']['countryCode']"
                                        }
                                        first_name      = {
                                            _copy = "submission.steps[4]['Address']['firstName']"
                                        }
                                        last_name       = {
                                            _copy = "submission.steps[4]['Address']['lastName']"
                                        }
                                        plot_area       = {
                                            _copy = "submission.steps[4]['Address']['plotArea']"
                                        }
                                        plot_of_land    = {
                                            _copy = "submission.steps[4]['Address']['plotOfLand']"
                                        }
                                        postal_code     = {
                                            _copy = "submission.steps[4]['Address']['zipCode']"
                                        }
                                        salutation      = {
                                            _copy = "submission.steps[4]['Address']['salutation']"
                                        }
                                        street          = {
                                            _copy = "submission.steps[4]['Address']['streetName']"
                                        }
                                        street_number   = {
                                            _copy = "submission.steps[4]['Address']['houseNumber']"
                                        }
                                        suburb          = {
                                            _copy = "submission.steps[4]['Address']['suburb']"
                                        }
                                        title           = {
                                            _copy = "submission.steps[4]['Address']['title']"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "country",
                                    "city",
                                    "postal_code",
                                    "street",
                                    "street_number",
                                    "additional_info",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "delivery_address"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[5]['Payment Method']"
                            }
                            origin    = "system_recommendation"
                            target    = "payment_method"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                "Copy of Text Input" = {
                                    _copy = "submission.steps[0]['Copy of Text Input']"
                                }
                                "Single Choice"      = {
                                    _copy = "submission.steps[2]['Single Choice']"
                                }
                                "Single Choice 2"    = {
                                    _copy = "submission.steps[2]['Single Choice 2']"
                                }
                                "Text Input"         = {
                                    _copy = "submission.steps[0]['Text Input']"
                                }
                            }
                            origin    = "system_recommendation"
                            target    = "journey_data"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _append = [
                                    {
                                        entity_id = {
                                            _copy = "submission.journey_context.relation_id"
                                        }
                                    },
                                ]
                                _uniq   = [
                                    "entity_id",
                                ]
                            }
                            origin    = "system_recommendation"
                            target    = "billing_contact.$relation"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['firstName']"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_first_name"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['lastName']"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_last_name"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['email']"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_email"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['telephone']"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_phone"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.steps[4]['Contact Info']['companyName']"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_company_name"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "account.tax_id"
                            }
                            origin    = "system_recommendation"
                            target    = "billing_vat"
                        }
                    )
                },
                {
                    any = jsonencode(
                        {
                            operation = {
                                _copy = "submission.source"
                            }
                            origin    = "system_recommendation"
                            target    = "source"
                        }
                    )
                },
            ]
            name                        = "Order from Journey"
            relation_attributes         = [
                {
                    mode                       = "set"
                    source_filter              = {
                        relation_tag = "customer"
                        self         = false
                    }
                    target                     = "billing_contact"
                    target_tags                = [
                        "customer",
                    ]
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
            target_schema               = "order"
            target_unique               = []
        },
    ]
}
