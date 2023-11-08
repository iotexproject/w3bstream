# w3bstream-mainnet
w3bstream mainnet

### step 1
install wsctl https://github.com/machinefi/w3bstream-mainnet/releases

### step 2
install docker & docker-compose

### step 3
```bash
git clone https://github.com/machinefi/w3bstream-mainnet.git
cd w3bstream-mainnet
```

### step 4
```bash
wsctl node up --private-key "your private key"
```

### step 5
```bash
wsctl node log
```

### step 6 
open a new terminal and execute
```bash
wsctl message send -p "test01" -v "0.1" -d "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}"
```
will see log on `wsctl node log` terminal

### step 7
```bash
wsctl node down
```

### step 8
can watch config by 
```bash
wsctl config get
```
and set config by 
```bash
wsctl config set
```