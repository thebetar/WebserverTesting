from flask import Flask
import json

app = Flask(__name__)

@app.get('/')
def get_pokemons():
    pokemon_str = open('../data/pokemon.json')
    pokemon = json.load(pokemon_str)

    parsed_pokemon = [
        { **item, 'legendary': item['Legendary'] == 'True' }
        for item in pokemon
    ]

    return parsed_pokemon