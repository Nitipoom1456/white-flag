-- +goose Up
ALTER TABLE feature_flag_setting
    DROP CONSTRAINT feature_flag_setting_min_app_version_id_fkey,
    DROP CONSTRAINT feature_flag_setting_max_app_version_id_fkey;

-- +goose Down
ALTER TABLE feature_flag_setting
    ADD CONSTRAINT feature_flag_setting_min_app_version_id_fkey FOREIGN KEY (min_app_version_id) REFERENCES app_version(uuid),
    ADD CONSTRAINT feature_flag_setting_max_app_version_id_fkey FOREIGN KEY (max_app_version_id) REFERENCES app_version(uuid);
