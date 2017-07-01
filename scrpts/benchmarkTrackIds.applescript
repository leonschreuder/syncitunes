# This is used to benchmark retreiving tracks with id info
tell application "iTunes"
	set resultList to {}
	repeat with currentPlaylist in (get every playlist)
		set playlistName to name of currentPlaylist
		set playlistID to id of currentPlaylist
		set parentID to -1
		try
			set parentID to id of parent in currentPlaylist
		end try
		set trackLocations to {}
		if class of currentPlaylist is user playlist then
			#Don't get tracks for folders
            # 1m17.266s
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
			#		copy {id, location} in currentTrack to end of trackLocations
			#	end try
			#end repeat

            # 0m3.301s -- Just repeat
			#repeat with currentTrack in (get every track in currentPlaylist)
			#end repeat

            # 0m1.507s --repeat with try
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
			#	end try
			#end repeat

            # 0m14.102s --repeat with retrieving values: 
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
			#		{id, location} in currentTrack
			#	end try
			#end repeat

            # 0m19.531s -- Other append syntax (same output as 'of every track')
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
            #        set trackLocations to trackLocations & ({id, location} in currentTrack)
			#	end try
			#end repeat

            # 0m20.076s
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
            #        set trackId to id in currentTrack
            #        set trackLocation to location in currentTrack
            #        set trackLocations to trackLocations & {trackId, trackLocation}
			#	end try
			#end repeat

            # 2m17.450s -- way slower with extra array creation
			# repeat with currentTrack in (get every track in currentPlaylist)
			# 	try
            #         set trackLocations to trackLocations & {{id, location} in currentTrack}
			# 	end try
			# end repeat

            # 0m18.274s -- output witout {} around each item
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
            #        set locationId to {id, location} in currentTrack
            #        set trackLocations to trackLocations & locationId
			#	end try
			#end repeat

            # 2m49.039s -- extra array,  same output but way slower
			#repeat with currentTrack in (get every track in currentPlaylist)
			#	try
            #        set locationId to {id, location} in currentTrack
            #        set trackLocations to trackLocations & {locationId}
			#	end try
			#end repeat

            # 0m2.674s -- results in separate lists with locations and ids
            # Will use this one and fix it in go
			try
				set trackLocations to {id, location} of every track in currentPlaylist
			end try

		end if
		set isSmart to false
		if class of currentPlaylist is not folder playlist then
			try
				set isSmart to smart of currentPlaylist
			end try
		end if
		if not isSmart then
            #performs worse than below for some reason
            #set resultList to resultList & {{playlistName, playlistID, parentID, trackLocations}}

			copy {playlistName, playlistID, parentID, trackLocations} to end of resultList
		end if
	end repeat
	return resultList
end tell
