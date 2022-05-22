# Herpetophobia

Service is just The Snake game.
On every level food spawns every 32 rounds and lives 3 rounds.
You must collect all 8 food to win and get the prize.
It is possible only if you know where the food will spawn.
Therefore you must hack the field generation algorythm and predict next field.

## Vuln

Field for the game generates as `field = init_permutation ^ (counter xor secret)`
We know only field and counter.

First of all, we need to gather some pairs of `(counter, field)`.
Then, using some Groups of Permutations magic we can calculate the `init_permutation` for certain level.
After that we calculate some suitable `secret` and can predict level field for any `counter`.

Thorough explanation of the vulnerability and exploitation is [here](./predict.ipynb)

After we figured out, what level field will go next, we must complete this level.
Here you can use any algorythm to find the best way.

Complete sploit is [here](./sploit.sage) 
