/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Spin, Image } from '@coze-arch/coze-design';

import { getIconByExtension, getFileExtension } from './utils';
import { FileItemStatus, PREVIEW_IMAGE_TYPE } from './constants';

import styles from './file-icon.module.less';

interface FileIconProps {
  file: {
    name: string;
    url?: string;
    status?: string;
  };
  size?: number;
}

export const FileIcon = (props: FileIconProps) => {
  const { size = 20, file } = props;

  const { url, name, status } = file;

  const extension = getFileExtension(name);

  if (status === FileItemStatus.Uploading) {
    return (
      <Spin
        wrapperClassName={styles['file-icon-loading']}
        style={{ width: size, height: size, lineHeight: `${size}px` }}
        spinning
      />
    );
  }

  if (PREVIEW_IMAGE_TYPE.includes(extension)) {
    return (
      <Image
        preview={false}
        className="object-contain object-center rounded-sm border-0"
        style={{ width: size, height: size }}
        imgStyle={{ width: size, height: size }}
        src={url}
        alt=""
      />
    );
  }

  const Icon = getIconByExtension(extension);

  return <Icon style={{ width: size, height: size, fontSize: size }} />;
};
