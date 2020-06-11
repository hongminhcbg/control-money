### control-money
## How to run this example
    step 1: create volume
        $ docker volume create controlmoney_db_data
    step 2: run two containers with docker compose
        $ docker-compose up
    step 3: data base migration
        - login mysql container with user: root, password: bW90aGVyIGZ1Y2tlciBub29i
        - create mysql table by copy file scripts/log.sql and scripts/user.sql, paste to mysql terminal

## API   
# /api/v1/profile
    bref: create user
    method: POST
    header auth:
        api-key:controlnomey-hongminh-229297
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
    header auth:
        api-key:controlnomey-hongminh-229297
    body:
        {
            "username": "xxx",
            "password":"xxx"            
        }
    
    response:
        {
            "data": {
                "id": 100,
                "username":"xxx",
                "tocken":"xxx",
                "name":"xxx",
                "money": 100                                
            },
            "meta_data": {
                "code": "200",
                "message":"success"
            }
        }

# /api/v1/log
    bref: create log 
    method: POST
    auth: tocken
    body:
        {
            "money": 123,
            "detail": "an trua",
            "tag":"abc"
        } 

# /api/v1/analysis/tag?begin=YYYY-MM-DD&end=YYYY-MM-DD
    bref: analysis monney depend on tag, param require Unix timestamp second
    method: GET
    auth: tocken
    response:
    {
        "data": [
            {
                "tag":"tag1",
                "money": 100
            },
            {
                "tag":"tag2",
                "money": 100
            },
            {
                "tag":"tag3",
                "money": 100
            },
            .
            .
            .
        ],
        "meta_data": {
            "code": "200",
            "message": "success"
        }
    }
    
# /api/v1/analysis/day?begin=YYYY-MM-DD&end=YYYY-MM-DD
    bref: analysis monney depend on day, param require Unix timestamp second
    method: GET
    auth: tocken
    response:
    {
        "data": [
            {
                "day": "yyyy-mm-dd",
                "money": 100
            },
            {
                "day": "yyyy-mm-dd",
                "money": 100
            },
            {
                "day": "yyyy-mm-dd",
                "money": 100
            },
            {
                "day": "yyyy-mm-dd",
                "money": 100
            },
            .
            .
            .
        ],
        "meta_data": {
            "code": "200",
            "message":"success"
        }        
    }

# /api/v1/average/day?begin=YYYY-MM-DD&end=YYYY-MM-DD
    bref: calc average monney spend per day
    method: GET
    param: 1 to 12
    auth: tocken
    response:
        {
            "data":{
                "money_spend_per_day": 100
            },
            "meta_data": {
                "code": "200",
                "message":"success"
            }
        }
## Tag is constant
    - xăng xe
    - thức ăn 
    - tiết kiệm
    - cho vay
    - quần áo
    - tiền nhà
    - đồ dùng, đồ vệ sinh cá nhân
    - đồ uống
    - khác
    