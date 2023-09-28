# Sudoku

**Sudoku**是一个基于`Vue2`和`Go`的数独游戏，可以自动生成数独题目，支持多种难度的数独题目。

## Project Background

jason哥最近迷上了九宫格数独，所谓九宫格数独是把一个9x9的网格从左到右从上到下依次划分出9个3x3的子网格，每个3x3的子网格称为一宫，共划分成9宫。而9宫数独是在网格的格子上填上1到9的数字，要求每一行每一列和每一宫都刚好填上1到9的数字不重复不遗漏.对于聪明的jason哥来说，只做一个九宫格数独实在是太简单了，她可以同时完成9个九宫格的数独求解，但每次玩数独游戏时，每一个九宫格数独都是串行生成的，暴躁的jason哥不愿意等待串行生成的时间，于是需要你并发生成9个九宫格数独，注意本次作业只需要并发生成9个数独，而不是求解，但注意所生成数独的可解性，具体的求解交给聪明的jason哥就可以啦。

## Quick Start

`前端` **sudoku文件**

```
npm install
```

```
npm run serve
```

## Project Structure

### Overall

```
├─.gitattributes
├─README.md
├─sudoku_golang     // go
├─sudoku            // frontend
└─.git
```

### Sudoku_frontend

```
├─sudoku
|   ├─.browserslistrc
|   ├─.gitignore
|   ├─.postcssrc.js
|   ├─babel.config.js
|   ├─jsconfig.json
|   ├─package-lock.json
|   ├─package.json
|   ├─README.md
|   ├─vue.config.js
|   ├─yarn.lock
|   ├─src
|   |  ├─App.vue
|   |  ├─main.js
|   |  ├─views
|   |  ├─utils
|   |  ├─router
|   |  ├─helper
|   |  ├─assets
|   |  ├─api
|   ├─public
└─.git
```

### Sudoku_golang

```
├─suduku_golang
|   ├─handle
|   ├─interface
|   ├─router
|   ├─service
|   ├─test
|   ├─utils
|   ├─src
|   ├─go.mod
|   ├─go.sum
|   ├─sudo.go
|   ├─sudo
└─.git
```

