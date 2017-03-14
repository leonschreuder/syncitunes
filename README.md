
Syncitunes syncs your local folder structure to itunes.


So it turns this
    
    music/artist/album/01song.mp3
    music/artist/album/02song.mp3
    music/other_artist/album/01song.mp3

into this

    ▾ music             <- folder
      ▾ artist          <- folder
        ▾ album         <- playlist
            01song.mp3  <- songs in playlist
            02song.mp3
      ▾ other_artist
        ▾ album
            song.mp3


The current feature set includes:

    - mac exclucivity
    - excruciating slowness
    - ignoring of what you already have in iTunes


I might possibly fix this.


# How to

Everything is hardcoded now, so just don't.
