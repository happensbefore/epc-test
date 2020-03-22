Тестовое задание для епк

Запуск приложения:

>cd docker

>docker-compose up

Сделать POST запрос на localhost:8080/product

пример тела запроса:
```json
{
    "product": {
        "name": "Игровой",
        "components": [
            {
                "isMain": true,
                "name": "Интернет",
                "prices": [
                    {
                        "cost": 100,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "adsl"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "10"
                            }
                        ]
                    },
                    {
                        "cost": 150,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "adsl"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "15"
                            }
                        ]
                    },
                    {
                        "cost": 500,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "xpon"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "100"
                            }
                        ]
                    },
                    {
                        "cost": 900,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "xpon"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "200"
                            }
                        ]
                    },
                    {
                        "cost": 200,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "fttb"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "30"
                            }
                        ]
                    },
                    {
                        "cost": 400,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "fttb"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "50"
                            }
                        ]
                    },
                    {
                        "cost": 600,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "fttb"
                            },
                            {
                                "codeName": "internetSpeed",
                                "operator": "EQ",
                                "value": "200"
                            }
                        ]
                    },
                    {
                        "cost": 10,
                        "priceType": "DISCOUNT",
                        "ruleApplicabilities": [
                            {
                                "codeName": "internetSpeed",
                                "operator": "GTE",
                                "value": "50"
                            }
                        ]
                    },
                    {
                        "cost": 15,
                        "priceType": "DISCOUNT",
                        "ruleApplicabilities": [
                            {
                                "codeName": "internetSpeed",
                                "operator": "GTE",
                                "value": "100"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "ADSL Модем",
                "prices": [
                    {
                        "cost": 300,
                        "priceType": "COST",
                        "ruleApplicabilities": [
                            {
                                "codeName": "technology",
                                "operator": "EQ",
                                "value": "adsl"
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "conditions": [
        {
            "ruleName": "technology",
            "value": "adsl"
        },
        {
            "ruleName": "internetSpeed",
            "value": "10"
        }
    ]
}
```
