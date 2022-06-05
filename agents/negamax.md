# NegaMax

NegaMax is a simplification of minimax for two-player, zero-sum games. It relies
on the property that one player's gain is equivalent to another's loss.

The pseudocode for the algorithm will be:

```
function negamax(node, depth, color) {
  if depth == targetDepth or node is a terminal:
    return color * heuristicValue(node, depth)
  value = −Infinity
  for each child of node:
    bestValue = max(bestValue, -negamax(child, depth + 1, -color))
  return bestValue
}
```
