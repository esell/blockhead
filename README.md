[![Build Status](https://drone.esheavyindustries.com/api/badges/esell/blockhead/status.svg)](https://drone.esheavyindustries.com/esell/blockhead)
[![Coverage](http://esheavyindustries.com:8080/display?repo=blockhead_git)](http://esheavyindustries.com:8080/display?repo=blockhead_git)

# Purpose

Blockhead is meant to be an interactive way to learn about blockchain technology. It is not designed to be a full blown, real deal blockchain, but
instead to allow users to interact with a toy blockchain and understand the basic concepts.

Blockhead acts as a standalone node where all of the typicaly blockchain processes are ran. There will be a built-in web UI where users
can add transactions, trigger mining, edit transactions etc in order to see, step-by-step, how a blockchain works.

Blockhead started as an attempt at re-implementing the simple chain shown [here](https://blockchain.works-hub.com/blog/Learn-Blockchains-by-Building-One) but has 
since evolved into (hopefully) a learning tool.


# Typical process

* Start the app :)
* Create a transaction: `curl -XPOST 'localhost:8000/newTransaction?to=blahto&from=blahfrom&amount=999'`
* Mine a new block: `curl 'localhost:8000/mine'`
* Get a list of blocks: `curl 'localhost:8000/list'`



