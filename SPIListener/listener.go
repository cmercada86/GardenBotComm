package listener

type (
	//SPIlistener listens on the SPI bus
	SPIListener struct {
		spi *SPIDevice
	}
)

func NewSPIListener() *SPIListener {

	return &SPIListener{}
}
