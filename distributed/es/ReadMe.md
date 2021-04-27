
```js
    /*


    比如我们有一段bool表达式，user_id = 1 and (product_id = 1 and (star_num = 4 or star_num = 5) and banned = 1)，写成SQL是如
    下形式：

    select * from xxx where user_id = 1 and (
    product_id = 1 and (star_num = 4 or star_num = 5) and banned = 1

    */
	

```
上面的表达式写成es的DSL是如下形式：
      
```json
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "user_id": {
              "query": "1",
              "type": "phrase"
            }
          }
        },
        {
          "match": {
            "product_id": {
              "query": "1",
              "type": "phrase"
            }
          }
        },
        {
          "bool": {
            "should": [
              {
                "match": {
                  "star_num": {
                    "query": "4",
                    "type": "phrase"
                  }
                }
              },
              {
                "match": {
                  "star_num": {
                    "query": "5",
                    "type": "phrase"
                  }
                }
              }
            ]
          }
        },
        {
          "match": {
            "banned": {
              "query": "1",
              "type": "phrase"
            }
          }
        }
      ]
    }
  },
  "from": 0,
  "size": 1
}
```