%function generate_astro_data(start_time_str, end_time_str, chunk_size, output_dir, output_version, sample_rate)
%    % Generate astro inertial data by time and store it in chunks as multiple .mat files
%    % Support Save Format Version: v5(-v6), v7(-v7), v7.3(-v7.3)
%
%    % ===== Default parameters =====
%    if nargin < 6
%        sample_rate = 10;  % default 10Hz
%    end
%    if nargin < 5
%        output_version = 'v5';
%    end
%    if nargin < 4
%        output_dir = 'testcase/output/';
%    end
%    if nargin < 3
%        chunk_size = 10000;
%    end
%
%    % concat output directory
%    output_dir = strcat(output_dir, '_', output_version);
%    if ~exist(output_dir, 'dir')
%        mkdir(output_dir);
%    end
%
%    % ===== Time settings =====
%    start_time = datetime(start_time_str, 'InputFormat', 'yyyy-MM-dd HH:mm:ss');
%    end_time   = datetime(end_time_str,   'InputFormat', 'yyyy-MM-dd HH:mm:ss');
%    dt = seconds(1 / sample_rate);  % set frequence
%
%    chunk_index = 1;
%    entry_index = 0;
%    chunk_data = [];  % generate entry
%    entry_template = []; % create null template
%
%    current_time = start_time;
%
%    while current_time <= end_time
%        % 序列化为时间字符串
%        % entry.timestamp     = current_time; % 如果不序列化 会存储为MCOS Matlab Class Object Serial
%        % entry.device_id     = "device_001";
%        % 序列化为时间戳
%        % entry.timestamp     = datestr(current_time, 'yyyy-mm-dd HH:MM:SS.FFF'); % 如果不序列化 会存储为MCOS Matlab Class Object Serial
%        entry.timestamp     = posixtime(current_time); % 如果不序列化 会存储为MCOS Matlab Class Object Serial
%        entry.device_id     = char("device_001"); % 如果不序列化 会存储为MCOS Matlab Class Object Serial
%        entry.status        = randi([0 7]);
%
%        entry.gx1 = (rand() - 0.5) * 2 * 1e6;
%        entry.gy1 = (rand() - 0.5) * 2 * 1e6;
%        entry.gz1 = (rand() - 0.5) * 2 * 1e6;
%        entry.ax1 = (rand() - 0.5) * 2000;
%        entry.ay1 = (rand() - 0.5) * 2000;
%        entry.az1 = (rand() - 0.5) * 2000;
%
%        entry.gx2 = (rand() - 0.5) * 2 * 1e6;
%        entry.gy2 = (rand() - 0.5) * 2 * 1e6;
%        entry.gz2 = (rand() - 0.5) * 2 * 1e6;
%        entry.ax2 = (rand() - 0.5) * 2000;
%        entry.ay2 = (rand() - 0.5) * 2000;
%        entry.az2 = (rand() - 0.5) * 2000;
%
%        entry.outer_angle   = rand() * 360 * 36000;
%        entry.inner_angle   = rand() * 360 * 36000;
%        entry.sampling_rate = randi([1 2]);
%        entry.pitch_angle   = (rand() - 0.5) * 20;
%        entry.roll_angle    = (rand() - 0.5) * 20;
%        entry.yaw_angle     = rand() * 360;
%        entry.vx            = (rand() - 0.5) * 200;
%        entry.vy            = (rand() - 0.5) * 200;
%        entry.vz            = (rand() - 0.5) * 20;
%        entry.longitude     = (rand() - 0.5) * 360;
%        entry.latitude      = (rand() - 0.5) * 180;
%
%        if entry_index == 0
%            chunk_data = repmat(entry, 0, 1);  % initial list structure
%            entry_template = entry;
%        end
%
%        entry_index = entry_index + 1;
%        chunk_data(end+1) = entry;
%
%        % 保存分块
%        if mod(entry_index, chunk_size) == 0
%            filename = fullfile(output_dir, sprintf('chunk_%d.mat', chunk_index));
%            switch output_version
%                case 'v5'
%                    save(filename, 'chunk_data', '-v6');
%                case 'v7'
%                    save(filename, 'chunk_data', '-v7');
%                case 'v7.3'
%                    save(filename, 'chunk_data', '-v7.3');
%                otherwise
%                    error("not support version: %s", output_version);
%            end
%            fprintf("Saved %s with %d records\n", filename, length(chunk_data));
%            chunk_index = chunk_index + 1;
%            entry_index = 0;
%            chunk_data = [];  % reset list empty
%        end
%
%        current_time = current_time + dt;
%    end
%
%    if entry_index > 0
%        filename = fullfile(output_dir, sprintf('chunk_%d_%s_1.mat', chunk_index, output_version));
%        switch output_version
%            case 'v5'
%                save(filename, 'chunk_data', '-v6');
%            case 'v7'
%                save(filename, 'chunk_data', '-v7');
%            case 'v7.3'
%                save(filename, 'chunk_data', '-v7.3');
%            otherwise
%                error("不支持的版本号: %s", output_version);
%        end
%        fprintf("Saved %s with %d records\n", filename, length(chunk_data));
%    end
%
%    % export fields to json
%    if ~isempty(entry_template)
%        field_list = fieldnames(entry_template);
%        fid = fopen(fullfile(output_dir, 'fields.json'), 'w');
%            fprintf(fid, '%s\n', jsonencode(field_list));
%            fclose(fid);
%    end
%
%end



function generate_astro_data()
    % 设置随机数种子为当前时间，确保每次执行随机数据不同
    rng('shuffle');

    % mock_matrix_data.m
    num_samples = 10; % 模拟数据行数
    num_columns = 18; % 18 列数据

    base_time = datetime(2020, 1, 1, 1, 1, 1);

    % 获取当前时间（即系统时间）
    current_time = datetime(2025, 1, 1, 1, 1, 1);

    % 生成一个 3000 行，18 列的随机数矩阵
    mock_data = rand(num_samples, num_columns);  % 随机数矩阵，可以替换为其他类型的数据

    % 生成时间序列：从当前时间加 1 秒开始
    for i = 1:num_samples
        % 当前时间 + 1 秒
        current_time_plus_1s = current_time + seconds(i);

        % 计算与基准时间的时间差（以秒为单位），并保留三位小数
        time_diff = seconds(current_time_plus_1s - base_time);
        time_diff = round(time_diff, 3);  % 保留三位小数

        % 将计算的时间差更新到第 7 列和第 8 列
        mock_data(i, 7) = time_diff;  % 第7列
        mock_data(i, 8) = time_diff;  % 第8列
    end

    % 保存为 v7.3 格式
    save('./testcase/mock_data.mat', 'mock_data', '-v7.3');
end

