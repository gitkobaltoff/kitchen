package main

const testingPayload = `"order_id": 1,
"table_id": 1,
"waiter_id": 1,
"items": [ 3, 4, 4, 2 ],
"priority": 3,
"max_wait": 45,
"pick_up_time": 1631453140 // UNIX timestamp
"cooking_time": 65
"cooking_details": [
{
"food_id": 3,
"cook_id": 1,
},
{
"food_id": 4,
"cook_id": 1,
},
{
"food_id": 4,
"cook_id": 2,
},
{
"food_id": 2,
"cook_id": 3,
},
]
`
