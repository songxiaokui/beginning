#include "stdio.h"
#include "string.h"
#include "extern.h"

#define PI 3.141592653
extern int AGE; // 只声明

static int count = 10;  // 整个文件可见

void f1()
{
    static int n = 5; // 只会在第一次调用初始化
    n++;
    printf("count = %d, n = %d\n", count, n);
}

const double Pi = 3.14;
int main(void)
{
    printf("hello day1 110 c...\n");
    printf("PI = %.3f\n", PI);
    printf("Pi = %.2f\n", Pi);

    {
        // auto 存储自动存储 分配在栈上 函数调用时创建 函数执行完毕释放内存
        auto char name[] = "sxk";
        printf("name is: %s\n", name);
        printf("name address is: %p\n", &name);

        // register 寄存器存储 没有内存地址
        register int a = 0;
        // printf("a address is %p\n", &a);  // 报错
    }

    {
        while (count--)
        {
            f1();
        }
    }

    {
        // extern 即可只声明 也可声明并初始化
        printf("extern AGE = %d\n", AGE);
    }
}
