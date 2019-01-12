//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeProposalCreate

package samples

func init() {

	sampleDataProposalCreateOperation[43] = `{
  "expiration_time": "2016-08-15T15:46:00",
  "extensions": [],
  "fee": {
    "amount": 2331876,
    "asset_id": "1.3.0"
  },
  "fee_paying_account": "1.2.119257",
  "proposed_ops": [
    {
      "op": [
        6,
        {
          "account": "1.2.119257",
          "active": {
            "account_auths": [],
            "address_auths": [],
            "key_auths": [
              [
                "BTS64hCxiD34egngzTHzTgCxpD5eazcQWxJEXoxGvrrag88epi2YD",
                1
              ]
            ],
            "weight_threshold": 1
          },
          "extensions": {},
          "fee": {
            "amount": 30928,
            "asset_id": "1.3.0"
          },
          "new_options": {
            "extensions": [],
            "memo_key": "BTS64hCxiD34egngzTHzTgCxpD5eazcQWxJEXoxGvrrag88epi2YD",
            "num_committee": 0,
            "num_witness": 0,
            "votes": [
              "2:148",
              "2:171",
              "2:179"
            ],
            "voting_account": "1.2.5"
          },
          "owner": {
            "account_auths": [
              [
                "1.2.119257",
                50
              ],
              [
                "1.2.119322",
                50
              ]
            ],
            "address_auths": [],
            "key_auths": [
              [
                "BTS6Mb4FyEqCtRxAZd5ihHTC1NZJn4AsTrZv3aD4uucUQJ33W5Lku",
                1
              ]
            ],
            "weight_threshold": 100
          }
        }
      ]
    }
  ]
}`

}

//end of file