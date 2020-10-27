# JavaScript

## [第一周](https://github.com/ncuhome/ncuhome-weekly-puzzle/blob/master/docs/js/w1.md)

输出：`undefined` 和 `ReferenceError`

解释：通过 `var` 和 `function` 声明的变量有**变量提升**，而 `let`、`const` 和 `class` 没有。

## [第五周](https://github.com/ncuhome/ncuhome-weekly-puzzle/blob/master/docs/js/w5.md)

输出：`undefined`、`2`、`3` 和 `5`

解释：

- `a = 1` 给全局变量 `window` 添加了属性 `a`，可以被 `delete` 删除
- 通过 `var`, `let` 和 `const` 创建的属性无法被 `delete` 删除
- `delete` 只会删除对象自身拥有的属性，不会删除原型链上的属性
- `in` 操作符的对象是键而非值
- `in` 对被 `delete` 删除的属性返回 `false`
