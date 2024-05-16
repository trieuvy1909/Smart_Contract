package services

// import (
// 	"net/http"
//     "crypto/ecdsa"
// 	"backend_go/entities"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"

// 	"context"
//     "math/big"
// 	"os"
//     "github.com/ethereum/go-ethereum/accounts/abi/bind"
//     "github.com/ethereum/go-ethereum/common"
//     "github.com/ethereum/go-ethereum/crypto"
//     "github.com/ethereum/go-ethereum/ethclient"
 
// )

// type SmartContractService struct {
// 	db *gorm.DB
// }

// func NewSmartContractService(db *gorm.DB) *SmartContractService {
// 	return &SmartContractService{db}
// }
// func (s *SmartContractService) GetValue() gin.HandlerFunc {
// 	return func(c *gin.Context){
// 		reps := entities.BaseResponse{}
// 		reps.Status = 0
// 		reps.Data = ""
// 		client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to connect to the Ethereum client"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

// 		instance, err := NewContract(address, client)
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to instantiate contract"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		value, err := instance.GetValue(&bind.CallOpts{})
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to call contract method"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

		
// 		reps.Message = "Thành công"
// 		reps.Data = value
// 		c.JSON(http.StatusOK,reps)
// 	}
// }

// func (s *SmartContractService) SetValue() gin.HandlerFunc { 
// 	return func(c *gin.Context) {

// 		reps := entities.BaseResponse{}
// 		reps.Status = 0
// 		reps.Data = ""

// 		newValue := "newValue"
// 		client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to connect to the Ethereum client"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to load private key"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		publicKey := privateKey.Public()
// 		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 		if !ok {
// 			reps.Status = 1
// 			reps.Message = "Failed to cast public key to ECDSA"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to retrieve account nonce"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		gasPrice, err := client.SuggestGasPrice(context.Background())
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to suggest gas price"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		auth := bind.NewKeyedTransactor(privateKey)
// 		auth.Nonce = big.NewInt(int64(nonce))
// 		auth.Value = big.NewInt(0)
// 		auth.GasLimit = uint64(300000)
// 		auth.GasPrice = gasPrice

// 		address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

// 		instance, err := NewContract(address, client)
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to instantiate contract"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}

// 		_, err = instance.SetValue(auth, newValue)
// 		if err != nil {
// 			reps.Status = 1
// 			reps.Message = "Failed to send transaction"
// 			c.JSON(http.StatusInternalServerError, reps)
// 			return
// 		}
// 		reps.Message = "Thành công"
// 		reps.Data = newValue
// 		c.JSON(http.StatusOK,reps)
// 	}
// }
// func NewContract(address common.Address, client *ethclient.Client) (*Contract, error) {
// 	contractABI := `"[
// 		{
// 			"constant": false,
// 			"inputs": [
// 				{
// 					"name": "_value",
// 					"type": "uint256"
// 				}
// 			],
// 			"name": "setValue",
// 			"outputs": [],
// 			"payable": false,
// 			"stateMutability": "nonpayable",
// 			"type": "function"
// 		},
// 		{
// 			"constant": true,
// 			"inputs": [],
// 			"name": "getValue",
// 			"outputs": [
// 				{
// 					"name": "",
// 					"type": "uint256"
// 				}
// 			],
// 			"payable": false,
// 			"stateMutability": "view",
// 			"type": "function"
// 		}
// 	]
// 	"`
// 	instance, err := NewContract(address, contractABI, client)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return instance, nil
// }
