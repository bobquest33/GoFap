GoFap
=====

A simple http video streamer and playlist generator

What is it?
===========

Do you own a seedbox or any form of a remote server where you keep your TV Shows, Movies, Porn? GoFap helps you generate a playlist for VLC or any other video player that supports .m3u playlists.

Videos are streamed using a http server that GoFap implements. Playlist are ordered by time so your latest files will be displayed on top.

How to run it?
==============

    go run main.go --path /home/me/myfiles

If it starts without errors, you should be able to enter the following url into your video player:

    http://<hostname>:8000/playlist
