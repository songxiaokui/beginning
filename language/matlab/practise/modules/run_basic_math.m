function run_basic_math()
    A = [1, 2; 3, 4];
    B = [1, 2; 2,1];

    % 矩阵乘法 A矩阵的行 = B矩阵的列
    disp("矩阵乘法: ")
    C = A * B;
    disp("A*B = ");
    disp(C);

    % 元素dot乘法 对应位置乘法
    D = A .* B;
    disp("A.*B = ");
    disp(D);

    %  矩阵转置 行变成列 列变成行 两次转置就是原矩阵
    disp("矩阵转置: ")
    disp("A^T = ");
    disp(transpose(A));
    disp("A^T^T = ");
    disp((A')');
end