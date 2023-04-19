from typing import Dict, Union
import random
import colorama
from matplotlib import pyplot as plt


TransMat = Dict[int, Dict[Union[int, None], float]]

class Markovian:
    STATE_COLORS = {
            0: colorama.Fore.RED,
            1: colorama.Fore.BLUE,
            2: colorama.Fore.GREEN,
        }

    def __init__(self, transmat: TransMat, value: int):
        self.state = value 
        self._transmat = transmat

    def __iter__(self):
        return self

    def __next__(self):
        mat = self._transmat[self.state]
        next_value = random.choices([*mat.keys()], weights = mat.values(), k = 1)[0]
        if next_value is None:
            raise StopIteration
        self.state = next_value
        return self.state 

    def __str__(self):
        return f"{self.STATE_COLORS[self.state]}Markovian({self.state}){colorama.Fore.RESET}"

if __name__ == "__main__":
    user_transmat = {
            0: {
                0: 0.9,
                1: 0.05,
                None: 0.05
            },
            1: {
                0: 0.3,
                1: 0.95,
                2: 0.2,
            },
            2: {
                1: 0.95,
                2: 0.05,
            }
        }
    N = 30 
    users = [[*Markovian(user_transmat, 0)] for _ in range(N)]
    fig,axes = plt.subplots(N,1, sharex = True, sharey = True)
    fig.set_figheight(10)
    for u,ax in zip(users,axes):
        ax.plot(u)

    plt.savefig("img.png")
