/**
 * Copyright 2020-present Facebook. All Rights Reserved.
 *
 * This program file is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the
 * Free Software Foundation; version 2 of the License.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program in a file named COPYING; if not, write to the
 * Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor,
 * Boston, MA 02110-1301 USA
 */

package devices

type FlashDevice interface {
	GetType() string
	GetSpecifier() string
	GetFilePath() string
	GetFileSize() uint64
	MmapRO() ([]byte, error)
	Munmap([]byte) error
	Validate() error
}

// takes in the specifier and returns the FlashDevice
type FlashDeviceFactory = func(string) (FlashDevice, error)

// maps from the type of the device to the factory function
// that gets & validates the information of the flash storage device
// populated by each storage device via the registerFlashDeviceFactory
// function
var FlashDeviceFactoryMap map[string]FlashDeviceFactory = map[string]FlashDeviceFactory{}

func registerFlashDeviceFactory(deviceType string, factory FlashDeviceFactory) {
	FlashDeviceFactoryMap[deviceType] = factory
}
