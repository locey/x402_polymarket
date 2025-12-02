// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.20;
import "./MarketFactory.sol";
import "./Market.sol";

contract Settlement {
    MarketFactory public immutable factory;
    address[] public proxyAddr; 
    //代理账户的金额
    mapping(address => uint256) public userBalance;
    //用户买入某结果的份额,这里记录用户列表，方便查询份额
    address[] public user; 
    event SettlementMarket(address indexed proxyAddr, uint256 balance);

    constructor(address factoryAddr){
        require(factoryAddr != address(0), "Settlement: invalid factoryAddr");
        factory = MarketFactory(factoryAddr);
    }

    function settlementMarket() external {
        //require(factory != address(0), "factory should not zero address");
        address[] memory marketsAddr = factory.getAllMarkets();
        require(marketsAddr.length > 0, "marketsAddr length is zero");

        Market market;
        
        for(uint256 i=0; i < marketsAddr.length; i++){
            market = Market(marketsAddr[i]);
            //mapping(address => mapping(uint256 => uint256)) storage shareBalance = market.shareBalance;
            Market.MarketStatus currentStatus = market.status();
            if(currentStatus == Market.MarketStatus.Locked && currentStatus != Market.MarketStatus.Settled){
                market.setMarketStatusSettled();
                for(uint256 j=0; j < user.length; j++){
                    uint256 amount = market.getShareBalance(user[j], market.finalResult());
                    userBalance[user[j]] = amount;
                    emit SettlementMarket(user[j], amount);
                }
            }
        }
    }

    function getBalance() external view returns (uint256) {
        return userBalance[msg.sender];
    }
}