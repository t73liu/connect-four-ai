## MiniMax

MiniMax is a backtracking algorithm that alternates between two players: a
maximizer and a minimizer. The maximizer wants the highest score possible while
the minimizer wants the lowest. The algorithm uses an evaluation function to
estimate the value of a given board state. MiniMax assumes that both parties
play optimally.

In order to get the fastest possible win, the score must take the depth into
account (i.e. subtract the depth from the maximizer score and add the depth to
the minimizer score).

One possible evaluation function is to count the number of connected rows,
columns and diagonals for the maximizer. For example, the following board will
have a value of 4 for O (i.e. 2 diagonals + 1 column + 1 row). The evaluation
function will need to handle partial boards if the depth does not reach a
terminal node.

|   |   |   |   |   |  |   |
|   |   |   | O |   |  |   |
|   |   | O | O |   |  |   |
|   |   | X | X |   |  |   |
|   |   | O | X | X |  |   |
|   | O | X | X | O |  |   |

The pseudocode for the algorithm will be:

```
function minimax(node, depth, maximizingPlayer) {
  if depth = targetDepth or node is a terminal then
    return heuristicValue(node, depth)
  if maximizingPlayer then
    bestValue = −Infinity
    for each child of node do
      bestValue = max(bestValue, minimax(child, depth + 1, FALSE))
    return bestValue
  else
    bestValue := +Infinity
    for each child of node do
      bestValue = min(bestValue, minimax(child, depth + 1, TRUE))
    return bestValue
}
```
