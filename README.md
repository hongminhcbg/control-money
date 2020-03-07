## control-money

# /api/v1/profile
    bref: create user
    method: POST
    body:
        {
            "username": "xxx",
            "password":"xxx",
            "Name":"xxx"
        }
    
    response:
        {
            "message":"success"
        }

# /api/v1/login
    bref: login to web
    method: POST
    body:
        {
            "username": "xxx",
            "password":"xxx"            
        }
    
    response:
        {
            "message":"success",
            "tocken":"xxx"
        }

# /api/v1/log
    bref: create log 
    method: POST
    auth: tocken
    body:
        {
            "monney": 123,
            "detail": "an trua",
            "tag":"abc"
        } 

# /api/v1/analysis/tag?begin=123&end=456
    bref: analysis monney depend on tag, param require Unix timestamp second
    method: GET
    auth: tocken
    response:
    {
        "tag1": 100,
        "tag2": 100,
        "tag3": 100,
        .
        .
        .
    }
    
# /api/v1/analysis/day?begin=123&end=456
    bref: analysis monney depend on day, param require Unix timestamp second
    method: GET
    auth: tocken
    response:
    {
        "yyyy-mm-dd": 100,
        "yyyy-mm-dd": 100,
        "yyyy-mm-dd": 100,
        .
        .
        .
    }

# /api/v1/average?month=1
    bref: cacl average monney spend per day
    method: GET
    param: 1 to 12
    auth: tocken
    response:
        {
            "average": 100
        }
    
    