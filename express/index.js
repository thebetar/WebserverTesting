import fs from 'fs';
import express from 'express';
import compression from 'compression'

const PORT = 3000;

const app = express();

app.get('/', (req, res) => {
  const pokemonStr = fs.readFileSync('../data/pokemon.json', 'utf8');
  const pokemon = JSON.parse(pokemonStr);

  const parsedPokemon = pokemon.map(item => ({
    ...item,
    legendary: pokemon.Legendary === 'True'
  }));

  res.status(200).json(parsedPokemon);
});

app.use(compression());

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});

