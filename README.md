# B+ Tree - GO Implementation
Data structure (dynamic, multilevel, bounded)- representation of sorted data. It's advantages over lists or arrays is that insertion as well as removal is extremely quick.
All the data in B+ trees are saved in leaves and the intermediate notes only serve as guides to find values (keys and tree pointers).

### Extra info
* Record : node or block
* Order : maximum number of keys in a record
* Root : starting record
* Maximum number of keys: equal record
* Minimum number of keys: equal 1/2 record
* Maximum number of keys: order^height
* Minimum number of keys: 2(order/2)^(height-1)
