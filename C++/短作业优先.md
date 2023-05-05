```
#include<stdio.h>
#include<string.h>
/*
 * 5
A B C D E
0 1 2 3 4
4 3 1 2 4
 */
struct job{
    char name;
    int arrive_time;
    int serve_time;             // 服务时间
    int turnaround_time;        // 周转时间
    int end_time;
    double weighted_turnaround_time;
};

void sort_by_arrive_time(int n, struct job *jobs);
void sort_by_serve_time(int k, int n, struct job *jobs);
void short_job_first(int n, struct job *jobs);
void sort(int n, struct job *jobs);
void result(int n, struct job *jobs);
int main(){
    int n;
    scanf("%d", &n);
    struct job jobs[n];
    for(int i = 0;i < n; i++){
        scanf(" %c", &jobs[i].name);
    }
    for(int i = 0;i < n; i++){
        scanf("%d", &jobs[i].arrive_time);
    }

    for(int i = 0;i < n; i++){
        scanf("%d", &jobs[i].serve_time);
    }
    short_job_first(n , jobs);
    sort(n, jobs);
    result(n, jobs);
    return 0;
}


void sort_by_arrive_time(int n, struct job *jobs){
    char temp;
    int min;
    for(int j = 0; j < n - 1; j++){
        for(int i = 0; i < n - j - 1; i++){
            if(jobs[i].arrive_time > jobs[i+1].arrive_time){

                min = jobs[i].arrive_time;
                jobs[i].arrive_time = jobs[i+1].arrive_time;
                jobs[i+1].arrive_time = min;

                min = jobs[i].serve_time;
                jobs[i].serve_time =jobs[i+1].serve_time;
                jobs[i+1].serve_time = min;

                temp = jobs[i].name;
                jobs[i].name = jobs[i+1].name;
                jobs[i+1].name = temp;
            }
        }
    }
}

void sort_by_serve_time(int k, int n, struct job *jobs) {
    int running_job_num = 0; // 用running_job_num统计等待作业i运行的作业个数
    for (int j = k; j < n; j++) {
        if (jobs[j].arrive_time < jobs[k - 1].end_time) {
            running_job_num++;
        }
    }
    int min;
    char temp;
    //按最短运行时间排序
    for (int j = k; j < k + running_job_num - 1; j++) {
        for (int i = k; i < k + running_job_num + k - j - 1; i++) {
            if (jobs[i].serve_time > jobs[i + 1].serve_time) {
                min = jobs[i].arrive_time;
                jobs[i].arrive_time = jobs[i + 1].arrive_time;
                jobs[i + 1].arrive_time = min;

                min = jobs[i].serve_time;
                jobs[i].serve_time = jobs[i + 1].serve_time;
                jobs[i + 1].serve_time = min;

                temp = jobs[i].name;
                jobs[i].name = jobs[i + 1].name;
                jobs[i + 1].name = temp;
            }
        }
    }
}

void short_job_first(int n, struct job *jobs){


    sort_by_arrive_time(n, jobs);

    jobs[0].end_time = jobs[0].arrive_time + jobs[0].serve_time;//结束时间=到达时间+服务时间
    jobs[0].turnaround_time = jobs[0].serve_time;//结束时间=到达时间+服务时间
    jobs[0].weighted_turnaround_time = jobs[0].turnaround_time * 1.0 / jobs[0].serve_time;//带权周转时间=周转时间/服务时间


    for(int i=1;i<n;i++) {
        sort_by_serve_time(i,n,jobs);
        for(int i = 0; i < n; i++){
            printf("%d ",jobs[i].serve_time);
        }
        if (jobs[i].arrive_time >
            jobs[i - 1].end_time) {                                         //第i个进程到达系统时，第i-1个进程已运行完毕
            jobs[i].end_time = jobs[i].arrive_time + jobs[i].serve_time;
            jobs[i].turnaround_time = jobs[i].serve_time;
        } else {
            jobs[i].end_time = jobs[i - 1].end_time + jobs[i].serve_time;
            jobs[i].turnaround_time = jobs[i].end_time - jobs[i].arrive_time;
        }
        jobs[i].weighted_turnaround_time = jobs[i].turnaround_time * 1.0 / jobs[i].serve_time;
    }
}

void sort(int n, struct job *jobs){
    char temp;
    int min;
    double weighted_time;
    for(int j = 0; j < n - 1; j++){
        for(int i = 0; i < n - j - 1; i++){
            if(jobs[i].arrive_time > jobs[i+1].arrive_time){

                min = jobs[i].arrive_time;
                jobs[i].arrive_time = jobs[i+1].arrive_time;
                jobs[i+1].arrive_time = min;

                min = jobs[i].serve_time;
                jobs[i].serve_time =jobs[i+1].serve_time;
                jobs[i+1].serve_time = min;

                min = jobs[i].end_time;
                jobs[i].end_time =jobs[i+1].end_time;
                jobs[i+1].end_time = min;

                min = jobs[i].turnaround_time;
                jobs[i].turnaround_time =jobs[i+1].turnaround_time;
                jobs[i+1].turnaround_time = min;

                weighted_time = jobs[i].weighted_turnaround_time;
                jobs[i].weighted_turnaround_time =jobs[i+1].weighted_turnaround_time;
                jobs[i+1].weighted_turnaround_time = weighted_time;

                temp = jobs[i].name;
                jobs[i].name = jobs[i+1].name;
                jobs[i+1].name = temp;
            }
        }
    }
}

void result(int n, struct job *jobs){
    printf("作 业 名:");
    for(int i = 0; i < n - 1; i++){
        printf("%c ", jobs[i].name);
    }
    printf("%c\n", jobs[n-1].name);

    printf("到达时间:");
    for(int i = 0; i < n - 1; i++){
        printf("%d ", jobs[i].arrive_time);
    }
    printf("%d\n", jobs[n - 1].arrive_time);

    printf("服务时间:");
    for(int i = 0; i < n - 1; i++){
        printf("%d ", jobs[i].serve_time);
    }
    printf("%d\n", jobs[n - 1].serve_time);

    printf("完成时间:");
    for(int i = 0; i < n - 1; i++){
        printf("%d ", jobs[i].end_time);
    }
    printf("%d\n", jobs[n - 1].end_time);

    printf("周转时间:");
    for(int i = 0; i < n - 1; i++){
        printf("%d ", jobs[i].turnaround_time);
    }
    printf("%d\n", jobs[n - 1].turnaround_time);

    printf("带权周转时间:");
    for(int i = 0; i < n - 1; i++){
        printf("%.2f ", jobs[i].weighted_turnaround_time);
    }
    printf("%.2f\n", jobs[n - 1].weighted_turnaround_time);
}
```

