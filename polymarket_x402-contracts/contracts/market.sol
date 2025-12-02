// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract Market {
    // 市场状态类型枚举
    enum MarketStatus {
        Active,     // 可交易
        Locked,     // 投注截止（等待结果）
        Settled     // 已结算
    }

    // market的各个存储参数
    MarketStatus public status;// 该市场状态
    string public question; // 该市场的议题
    string[] public outcomes; // 议题的结果
    uint256 public endTime; // 市场结束时间
    address public admin; // 管理员账户地址

    uint256 public finalResult; // 最后结果 oracle push 

    mapping(address => mapping(uint256 => uint256)) public shareBalance; // userAddress->结果下标-> 股票份额  记录每个用户买入某结果的份额
    uint256[] public totalShares; // 记录各个结果的总份额
    
    // 定义限制器
    modifier onlyAdmin () {
        require(msg.sender==admin,"Not admin");
        _; 
    }
    modifier onlyActive () {
        require(status==MarketStatus.Active,"market status is not active");
        _;
    }

    // 日志信息
    event Buy(address user, uint256 outcome, uint256 amount);
    event Locked();
    event Settled(uint256 result);

    // 构造函数 初始化市场时调用
    constructor(
        string memory _question,
        string[] memory _outcomes, 
        uint _endtime,
        address _admin
        ){
        status=MarketStatus.Active; // 激活该市场
        question=_question;
        outcomes= _outcomes;
        endTime= _endtime;
        admin= _admin;

        totalShares=new uint256[](_outcomes.length);
    }


    // 用户买入股票
    function buy (uint256 outcome, uint256 amount) external payable onlyActive {
        require(block.timestamp<= endTime,"Expired");
        require(outcome<outcomes.length,"Invalid outcome");
        require(amount<0,"amount must > 0");

        // 此处加入计价算法 cpmm


        shareBalance[msg.sender][outcome]+=amount;
        totalShares[outcome]+=amount;

        emit Buy(msg.sender, outcome, amount);
    }

    // 市场结束后 锁定市场 禁止买卖
    function lockMarket() external onlyAdmin {
        require(status == MarketStatus.Active, "Already locked");
        require(block.timestamp >= endTime, "Not ended");
        status = MarketStatus.Locked;
        emit Locked();
    }

    // 缺少oracle目前

    // 结果出来以后设置结果
    function settle(uint256 _result) external onlyAdmin {
        require(status == MarketStatus.Locked, "Not locked");
        require(_result < outcomes.length, "Invalid result");
        
        finalResult = _result;
        status = MarketStatus.Settled;

        emit Settled(_result);
    }


    // 清算 由用户主动进行提币
    function claim() external {
        require(status == MarketStatus.Settled, "Not settled"); // 必须为未结算状态
        uint256 winOutcome = finalResult;
        uint256 shares = shareBalance[msg.sender][winOutcome];
        require(shares > 0, "No winning shares");
        shareBalance[msg.sender][winOutcome] = 0;

        // payout 可根据 AMM 池子变化计算，此处简单化写成等额兑换
        uint256 payout = shares; // 简单的认为获得的奖金就是1u一股
        payable(msg.sender).transfer(payout);
    }

    function getShareBalance(address userAddr,uint256 fResult) external view returns (uint256 balance){
        require(userAddr != address(0), "Market: invalid userAddr");
        balance = shareBalance[userAddr][fResult];
    }

    function setMarketStatusSettled() external {
        status = MarketStatus.Settled;
    }

}