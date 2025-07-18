# 概要

https://www.youtube.com/watch?v=VchfS2fvQCw

## 背景

バスケットボールでは、ショット精度以上にショットクリエイトが重要視される。特に NBA では、得点効率を意識したオフェンスが盛んであり、チームオフェンスのほとんどをエースが担うようなことも珍しくない。ところがエースが調子が悪い時には、そのセレクションに疑問が持たれることもある。そういうとき、SNS 上では、「外しても打ち続けるべき」又は「少しはセレクトを減らすべき」という意見が完全に 2 分される。
<br>
特にエースの象徴として Kobe Bryant の名が上がることが多い。ショットを外しても撃ち続ける姿勢はマンバメンタリティ(曖昧さ回避)と称えられる。

## 目的

マンバメンタリティがどうチームオフェンスに影響するのかを数理的な視点で理解することを目指す。
<br>
数理生物学的な手法で、バスケットボールにおけるショットセレクションモデルを理論上に構築し、シミュレーションを行うことで、これを目指す。

# モデル

＊工事中
描き途中な上に、いくつか式を変えてしまっています。

## 結果によるスコア成功率の変化

```math
s_p(n+1) = s_p^{base} + \eta_{p} \frac{1}{1+e^{-(n-n_{d})}} (acc(n) - s_p^{base})
```

```math
acc(n) = \frac{n_{made}(n)}{n}
```

ただし、$n$は試投数、$n_{made}(n)$は$n$時点での成功数を表す。
この時実際に減少が起こる確率は、以下のように定義される

```math
R(made) \coloneqq \left\{
\begin{array}{ll}
s_p(n) & (made=ture) \\
(1-s_p(n)) & (other) \\
\end{array}
\right.
```

```math
rate_{real}(n) = \prod_{i}^{n}{R(made(i))}
```

グラフは以下のようになった。

## 結果によるメンタルの変動

メンタル$mt$の変化量を以下にまとめる

```math
\frac{d mt}{dt} = M(t, shot, made)
```

```math
M(t, shot, made) \coloneqq \left\{
\begin{array}{ll}
0 & (shot = false) \\
\eta s_p(n) & (shot = true \land made = true) \\
-\eta (1-mamba) & (other) \\
\end{array}
\right.
```

## ショット選択

```math
P(i) = \frac{sp_{i}(t)}{OT + \sum_j{sp_{j}(t)}}
```

何も選択されなかった場合($OT$)は、その時点での最低値に依存するオフェンスを行う

```math
o_t \min(sp_i)
```

その際に選手メンタルを以下で更新する

```math
\frac{d mt}{dt} = - \eta o_t  (1-mamba)
```

# 実行

```bash
cd mamba
go run play/main.go
```

# TODO

- [ ] finish model
- [ ] finish chart tool
- [ ] web app
