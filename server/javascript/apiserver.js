const express = require('express');
const api = require('novelcovid');
const bodyParser = require('body-parser');
const app = express();

const mychannel = 'myemr';
const mycontract = 'emr';
const port = process.env.PORT || 5000

app.use(bodyParser.json());// to support JSON-encoded bodies
app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(express.static("public"));
app.set("view engine","ejs");
app.set("views","./views");

app.listen(port, async() => {
  console.log('RESTful API server started at: http://localhost:'+port)
});


const { FileSystemWallet, Gateway } = require('fabric-network');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', '..', 'network', 'connection-org1.json');

app.get('/covid', async (req, res) => {
  const global = await api.all();
  const ketquas = await api.countries({ sort: 'cases' });
  console.log(`Result is: $ketquas`);
  res.render('covid', { global, ketquas });
});

app.get('/api/queryallbn', async function (req, res) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork(mychannel);

        const contract = network.getContract(mycontract);

        const result = await contract.evaluateTransaction('queryAllBN');
        //console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        //res.status(200).json({response: result.toString()});
        //console.log(`Transaction has been evaluated, result is: ${response}`);
        res.status(200).json(JSON.parse(result.toString()));
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({error: error});
        //process.exit(1);
    }
});

app.get('/benhnhan/list', async function (req, res) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork(mychannel);

        const contract = network.getContract(mycontract);

        // string json
        let result = await contract.evaluateTransaction('queryAllBN');
        const ketqua = result.toString();
        //var ketqua = JSON.parse(response.toString());

        // Convert string to object
        const obj = JSON.parse(ketqua);
        //console.log(obj.Record[0]);
        //console.log(`Result is la: ${ketqua}`);

        //res.json(JSON.parse(response.toString()));
        res.render("dsbenhnhan.ejs",{obj})

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({error: error});
        //process.exit(1);
    }
});


app.get('/api/query/:bn_index', async function (req, res) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork(mychannel);
        const contract = network.getContract(mycontract);

        const result = await contract.evaluateTransaction('queryBN', req.params.bn_index);
        //console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        //res.status(200).json(result.toString());
        res.status(200).json(JSON.parse(result.toString()));
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({error: error});
        //process.exit(1);
    }
});

app.post('/api/addbn/', async function (req, res) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork(mychannel);
        const contract = network.getContract(mycontract);

        await contract.submitTransaction('createBN', req.body.bnid, req.body.mabn, req.body.hoten, req.body.ngaysinh, req.body.gioitinh, req.body.cmnd, req.body.diachi, req.body.maxa);
        console.log('Transaction has been submitted');
        res.send('Transaction has been submitted');
        await gateway.disconnect();
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        //process.exit(1);
    }
})

app.put('/api/changebn/:bn_index', async function (req, res) {
    try {
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork(mychannel);
        const contract = network.getContract(mycontract);

        await contract.submitTransaction('changeBN', req.params.bn_index, req.body.hoten);
        console.log('Transaction has been submitted');
        res.send('{"message": "Transaction has been submitted"}');
        await gateway.disconnect();
} catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        //process.exit(1);
    }
})

app.get("/",function(req,res){
  res.render("main");
});
