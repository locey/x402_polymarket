// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./Market.sol";

contract MarketFactory{
    address public owner; // 合约拥有者 后端管理员
    address[] public markets; // 市场列表
    mapping(address => bool) public isMarket; // 标记为是否为合法市场

    modifier onlyAdmin() {
        require(msg.sender==owner,"Not admin");
        _;
    }

    event MarketCreated(address marketAddress, string question);

    constructor(){
        owner= msg.sender;
    }

    // 创建市场
    function createMarket(
        string memory _question,string[] memory _outcomes,uint256 _endTime)external onlyAdmin returns (address){
            Market makrket= new Market(
                _question,
                _outcomes,
                _endTime,
                msg.sender // admin of market
            ); 
            markets.push(address(makrket));
            isMarket[address(makrket)]= true;
            emit MarketCreated(address(makrket), _question);
            return address(makrket);
    }

    // 获取所有市场 
    function getAllMarkets() external view returns (address[] memory) {
        return markets;
    }
}