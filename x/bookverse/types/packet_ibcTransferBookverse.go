package types

// ValidateBasic is used for validating the packet
func (p IbcTransferBookversePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcTransferBookversePacketData) GetBytes() ([]byte, error) {
	var modulePacket BookversesPacketData

	modulePacket.Packet = &BookversesPacketData_IbcTransferBookversePacket{&p}

	return modulePacket.Marshal()
}
