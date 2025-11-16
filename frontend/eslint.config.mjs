import { defineConfig, globalIgnores } from 'eslint/config';
import nextVitals from 'eslint-config-next/core-web-vitals';
import nextTs from 'eslint-config-next/typescript';
import importPlugin from 'eslint-plugin-import';
import simpleImportSort from 'eslint-plugin-simple-import-sort';

const eslintConfig = defineConfig([
  ...nextVitals,
  ...nextTs,
  // Override default ignores of eslint-config-next.
  globalIgnores([
    // Default ignores of eslint-config-next:
    '.next/**',
    'out/**',
    'build/**',
    'next-env.d.ts',
  ]),
  {
    plugins: {
      import: importPlugin,
      'simple-import-sort': simpleImportSort,
    },
    settings: {
      'import/resolver': {
        typescript: {
          alwaysTryTypes: true,
          project: './tsconfig.json',
        },
        node: true,
      },
    },
    rules: {
      // 使用 simple-import-sort 自动排序 import
      'simple-import-sort/imports': [
        'warn', // 编码时显示警告，commit时自动排序
        {
          groups: [
            // Node.js 内置模块
            ['^node:'],
            // React 和 Next.js
            ['^react', '^next'],
            // 第三方库
            ['^@?\\w'],
            // 内部模块（@/ 别名）
            ['^@/'],
            // 相对路径的父级目录
            ['^\\.\\.(?!/?$)', '^\\.\\./?$'],
            // 相对路径的同级目录
            ['^\\./(?=.*/)(?!/?$)', '^\\.(?!/?$)', '^\\./?$'],
            // 样式文件
            ['^.+\\.s?css$'],
            // 类型导入
            ['^.+\\u0000$'],
          ],
        },
      ],
      'simple-import-sort/exports': 'error',
      // 禁止重复导入
      'import/no-duplicates': 'error',
      // 确保导入的模块可以解析
      'import/no-unresolved': 'off', // Next.js 有自己的解析方式，关闭此规则
      // 优先使用 default export
      'import/prefer-default-export': 'off',
      // 禁止使用默认导出
      'import/no-default-export': 'off', // Next.js 需要 default export
    },
  },
]);

export default eslintConfig;
