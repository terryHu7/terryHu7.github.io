`
	{
		"Name": "metis_live_mcs_zlb",
		"Type": "MCS",
		"Description": "Pipeline task for metis live",
		"Version": "0.3",
		"CodecStreamSpecs": [
			{
				"Name": "Exp_front",
				"URI": "file:///home/user/Videos/eswin/trim0-118_30fps_1080p.mp4",
				"VideoCodec": "h.264"
			},
			{
				"Name": "Exp_top",
				"URI": "file:///home/user/Videos/eswin/trim1.24-119_30fps_1080p.mp4",
				"VideoCodec": "h.264"
			}
		],
		"ExperimentSpecs": [
			{
				"Name": "test",
				"DeskId": "123456",
				"TopViewStream": "Exp_top",
				"FrontViewStream": "Exp_front",
				"Code": "02010001",
				"Uuid": "7896661",
				"Step": "0201000101",
				"Width": 1920,
				"Height": 1080,
				"SnapshotPath": "/home/user/intelligent_experiment/7896661"
			}
		]
	}