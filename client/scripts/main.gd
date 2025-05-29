extends Node

const packets := preload("res://packets.gd")

func _ready() -> void:
	var data := [8, 12, 18, 14, 10, 12, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100]
	var packet := packets.Packet.new()
	packet.from_bytes(data)
	print(packet)
