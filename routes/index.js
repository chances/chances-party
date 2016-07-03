import express from 'express';

import {playlists, searchTracks} from '../lib/spotify';

let router = express.Router();

/* GET home page. */
router.get('/', function (req, res) {
  let pageData = {
    base: process.env.BASE_URL || '',
    user: req.user,
    playlists: undefined,
    error: undefined
  };
  if (req.user !== undefined && process.env.SPOTIFY_ACCESS_TOKEN) {
    req.user.currentPlaylist = null;
    playlists({
      id: req.user.id,
      token: process.env.SPOTIFY_ACCESS_TOKEN
    }, function (err, data) {
      if (err) {
        pageData.error = err;
        res.render('index', pageData);
      } else if (data.hasOwnProperty('error')) {
        pageData.error = data;
        res.render('index', pageData);
      } else {
        pageData.playlists = data.items;
        res.render('index', pageData);
      }
    });
  } else {
    res.render('index', pageData);
  }
});

router.get('/history', function (req, res, next) {
  // Get playback history (previously played songs)
  res.send('respond with a resource');
});

router.get('/queue', function (req, res, next) {
  // Get playback future (current track first?)
  // TODO: Maybe use streaming stuffs for current track/queue and history => Research
  res.send('respond with a resource');
});

router.get('/search', function (req, res, next) {
  let query = req.query.q || req.query.query || null;
  // let page = req.params.page || 0;
  if (query == null) {
    res.statusCode = 400;
    return res.json({message: 'Error 400: Malformed search, missing query'});
  }
  searchTracks({
    query: query
  }, function (err, data) {
    if (err) {
      res.statusCode = 500;
      return res.json({message: 'Error 500: ' + err.message});
    }

    res.json(data.tracks);
  });
});

export default router;
