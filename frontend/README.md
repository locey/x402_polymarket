# X402 Prediction Market Frontend

## 项目背景

X402 Prediction Market 是一个基于区块链的去中心化预测市场平台。该项目使用 Next.js 构建前端应用，集成了 RainbowKit 和 Wagmi 来实现 Web3 钱包连接和区块链交互功能。

### 技术栈

- **框架**: Next.js 16.0.3 (App Router)
- **语言**: TypeScript 5
- **UI 库**: React 19.2.0
- **Web3 集成**:
  - [RainbowKit](https://rainbowkit.com/) - 钱包连接 UI 组件库
  - [Wagmi](https://wagmi.sh/) - React Hooks 用于以太坊
  - [Viem](https://viem.sh/) - 类型安全的以太坊工具库
- **样式**: Tailwind CSS 4
- **数据获取与缓存**: TanStack Query (React Query) - 用于服务器状态管理和数据同步
- **代码质量**:
  - ESLint - 代码检查
  - Prettier - 代码格式化
  - Husky - Git hooks
  - lint-staged - 提交前代码检查

### 支持的区块链网络

- **主网**: Ethereum Mainnet, Polygon, Optimism, Arbitrum, Base
- **测试网** (可选): Sepolia, Polygon Amoy

## 开发规范

### 代码风格

1. **TypeScript**: 使用严格模式，所有代码必须通过类型检查
2. **命名规范**:
   - 组件文件使用 PascalCase: `BetCard.tsx`
   - 工具函数使用 camelCase: `formatPrice.ts`
   - 常量使用 UPPER_SNAKE_CASE: `MAX_BET_AMOUNT`
3. **文件组织**:
   - 页面组件放在 `src/app/` 目录
   - 可复用组件放在 `src/components/` 目录
   - 工具函数放在 `src/utils/` 目录（如需要）
   - 类型定义放在 `src/types/` 目录（如需要）
4. **路径别名**: 使用 `@/` 作为 `src/` 的别名

### 提交规范

项目使用 Husky 和 lint-staged 确保代码质量：

- 提交前会自动运行 ESLint 和 Prettier
- 只有通过检查的代码才能提交
- 提交信息应清晰描述变更内容

### 组件开发规范

1. **客户端组件**: 使用 Web3 hooks 或需要交互的组件必须添加 `'use client'` 指令
2. **服务端组件**: 默认使用服务端组件，除非需要客户端功能
3. **Provider 组件**: Web3 相关 Provider 应放在 `src/components/Layouts/` 目录

## 开发流程

### 环境要求

- Node.js 18+
- npm 或 yarn 包管理器

### 安装依赖

```bash
npm install
```

### 环境变量配置

创建 `.env 或 .env.local` 文件（如果不存在），配置以下变量：

```env
# WalletConnect Project ID (从 https://cloud.walletconnect.com/ 获取)
NEXT_PUBLIC_PROJECT_ID=your_project_id

# 是否启用测试网 (可选，默认 false)
NEXT_PUBLIC_ENABLE_TESTNETS=false
```

### 开发命令

#### 启动开发服务器

```bash
npm run dev
```

开发服务器将在 `http://localhost:3000` 启动。

#### 构建生产版本

```bash
npm run build
```

#### 启动生产服务器

```bash
npm run start
```

### 代码质量命令

#### 运行 ESLint 检查

```bash
npm run lint
```

#### 自动修复 ESLint 问题

```bash
npm run lint:fix
```

#### 格式化代码

```bash
npm run format
```

#### 检查代码格式

```bash
npm run format:check
```

### 项目结构

```
frontend/
├── src/
│   ├── app/                    # Next.js App Router 页面
│   │   ├── layout.tsx         # 根布局
│   │   ├── page.tsx           # 首页
│   │   └── market/            # 市场相关页面
│   │       └── [id]/          # 动态路由：市场详情页
│   ├── components/            # React 组件
│   │   ├── Cards/            # 卡片组件
│   │   ├── Layouts/          # 布局组件（Provider 等）
│   │   ├── Header.tsx        # 头部组件
│   │   ├── Footer.tsx        # 底部组件
│   │   └── SearchInput.tsx   # 搜索输入组件
│   └── wagmi.ts              # Wagmi 配置
├── public/                    # 静态资源
├── abis/                       # 合约 ABI 文件
├── next.config.ts            # Next.js 配置
├── tsconfig.json             # TypeScript 配置
├── eslint.config.mjs         # ESLint 配置
├── postcss.config.mjs        # PostCSS 配置
└── package.json              # 项目依赖和脚本
```

## 开发注意事项

1. **Web3 配置**: 在 `src/wagmi.ts` 中配置支持的区块链网络和 WalletConnect Project ID
2. **性能优化**: Next.js 配置中已启用包导入优化，减少 bundle 大小
3. **类型安全**: 充分利用 TypeScript 的类型系统，避免使用 `any` 类型
4. **错误处理**: Web3 操作应包含适当的错误处理和用户提示
5. **响应式设计**: 使用 Tailwind CSS 确保组件在不同设备上正常显示

## 常见问题

### WalletConnect Project ID

如果遇到钱包连接问题，请确保：

1. 在 [WalletConnect Cloud](https://cloud.walletconnect.com/) 注册并获取 Project ID
2. 在 `.env` 或 `.env.local` 中正确配置 `NEXT_PUBLIC_PROJECT_ID`

### 测试网支持

要启用测试网，在 `.env 或 .env.local` 中设置：

```env
NEXT_PUBLIC_ENABLE_TESTNETS=true
```

确保所有代码通过 ESLint 和 Prettier 检查后再提交。

## 技术说明

### TanStack Query (React Query) 简介

**TanStack Query**（原名 React Query）是一个专门用于**服务器状态管理**的库。

#### 主要功能

1. **数据获取与缓存**: 自动缓存从服务器（或区块链）获取的数据，避免重复请求
2. **后台同步**: 自动在后台刷新过期数据，保持数据最新
3. **请求去重**: 多个组件请求相同数据时，只发送一次请求
4. **加载与错误状态**: 提供 `isLoading`、`isError` 等状态，简化 UI 处理
5. **乐观更新**: 支持在数据提交前更新 UI，提升用户体验

#### 与传统状态管理的区别

| 特性         | TanStack Query               | Redux/Zustand                   |
| ------------ | ---------------------------- | ------------------------------- |
| **用途**     | 服务器状态（API/区块链数据） | 客户端状态（UI 状态、表单数据） |
| **数据来源** | 外部数据源（API、区块链）    | 应用内部状态                    |
| **缓存**     | 自动缓存和失效策略           | 需要手动管理                    |
| **同步**     | 自动后台同步                 | 需要手动触发更新                |

#### 在项目中的使用

在本项目中，TanStack Query 主要用于：

1. **Wagmi 内部使用**: Wagmi 库内部使用 TanStack Query 来管理区块链数据的获取和缓存
   - 账户余额查询
   - 交易历史
   - 合约状态读取
   - 事件监听

2. **数据同步**: 当多个组件需要相同的区块链数据时，TanStack Query 确保：
   - 只发送一次 RPC 请求
   - 数据在所有组件间共享
   - 自动处理加载和错误状态

#### 示例用法

虽然 Wagmi 已经封装了大部分功能，但如果你需要直接使用 TanStack Query，可以这样：

```typescript
import { useQuery } from '@tanstack/react-query';

// 获取数据
const { data, isLoading, error } = useQuery({
  queryKey: ['market', marketId],
  queryFn: () => fetchMarketData(marketId),
  staleTime: 1000 * 60 * 5, // 5分钟内数据视为新鲜
});

// 提交数据
const mutation = useMutation({
  mutationFn: newData => updateMarket(newData),
  onSuccess: () => {
    // 成功后刷新相关查询
    queryClient.invalidateQueries({ queryKey: ['market'] });
  },
});
```

#### 总结

TanStack Query 主要用于管理服务器状态。
对于客户端状态管理，如果需要，可以考虑使用：

- **React Context** - 简单状态共享
- **自定义Hooks/Models** - 用于管理客户端状态
