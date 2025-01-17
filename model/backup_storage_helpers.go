// percona-everest-backend
// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

// CreateBackupStorageParams parameters for BackupStorage record creation.
type CreateBackupStorageParams struct {
	Name        string
	Description string
	Type        string
	BucketName  string
	URL         string
	Region      string
	AccessKeyID string
	SecretKeyID string
}

// UpdateBackupStorageParams parameters for BackupStorage record update.
type UpdateBackupStorageParams struct {
	Name        string
	Description *string
	BucketName  *string
	URL         *string
	Region      *string
	AccessKeyID *string
	SecretKeyID *string
}

// CreateBackupStorage creates a BackupStorage record.
func (db *Database) CreateBackupStorage(_ context.Context, params CreateBackupStorageParams) (*BackupStorage, error) {
	s := &BackupStorage{
		Name:        params.Name,
		Description: params.Description,
		Type:        params.Type,
		BucketName:  params.BucketName,
		URL:         params.URL,
		Region:      params.Region,
		AccessKeyID: params.AccessKeyID,
		SecretKeyID: params.SecretKeyID,
	}
	err := db.gormDB.Create(s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ListBackupStorages returns all available BackupStorages records.
func (db *Database) ListBackupStorages(_ context.Context) ([]BackupStorage, error) {
	var storages []BackupStorage
	err := db.gormDB.Find(&storages).Error
	if err != nil {
		return nil, err
	}
	return storages, nil
}

// GetBackupStorage returns BackupStorage record by its Name.
func (db *Database) GetBackupStorage(_ context.Context, tx *gorm.DB, name string) (*BackupStorage, error) {
	gormDB := db.gormDB
	if tx != nil {
		gormDB = tx
	}
	storage := &BackupStorage{}
	err := gormDB.First(storage, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return storage, nil
}

// UpdateBackupStorage updates a BackupStorage record.
func (db *Database) UpdateBackupStorage(_ context.Context, tx *gorm.DB, params UpdateBackupStorageParams) error {
	target := db.gormDB
	if tx != nil {
		target = tx
	}
	old := &BackupStorage{}
	err := target.First(old, "name = ?", params.Name).Error
	if err != nil {
		return err
	}

	record := BackupStorage{}
	if params.Description != nil {
		record.Description = *params.Description
	}

	if params.BucketName != nil {
		record.BucketName = *params.BucketName
	}
	if params.URL != nil {
		record.URL = *params.URL
	}
	if params.Region != nil {
		record.Region = *params.Region
	}
	if params.AccessKeyID != nil {
		record.AccessKeyID = *params.AccessKeyID
	}
	if params.SecretKeyID != nil {
		record.SecretKeyID = *params.SecretKeyID
	}

	// Updates only non-empty fields defined in record
	if err = target.Model(old).Where("name = ?", params.Name).Updates(record).Error; err != nil {
		return errors.Join(err, errors.New("could not update backup storage"))
	}

	return nil
}

// DeleteBackupStorage returns BackupStorage record by its Name.
func (db *Database) DeleteBackupStorage(_ context.Context, name string, tx *gorm.DB) error {
	gormDB := db.gormDB
	if tx != nil {
		gormDB = tx
	}
	return gormDB.Delete(&BackupStorage{}, "name = ?", name).Error
}
