//
// Created by zlji on 2023/5/29.
//

#ifndef OPENMP_OPENMP_H
#define OPENMP_OPENMP_H

#include <stdio.h>
#include "omp.h"

void openmp_for();

void openmp_sections();

void openmp_single();

void openmp_for(){
    int NUM_THREADS = omp_get_num_procs();
#pragma omp parallel for
    for(int i = 0; i < NUM_THREADS; i++){
        printf("hello world! from thread_num %d\n", omp_get_thread_num());
    }
}

void openmp_sections(){
#pragma omp parallel sections
    {
#pragma omp section
        printf("section 1 from thread_num %d\n", omp_get_thread_num());
#pragma omp section
        printf("section 2 from thread_num %d\n", omp_get_thread_num());
#pragma omp section
        printf("section 3 from thread_num %d\n", omp_get_thread_num());
    }
}

void openmp_single(){
    omp_set_num_threads(4);
#pragma omp parallel
    {
#pragma omp single
        printf("Beijing work1. \n");
        printf("work on 1 parallellly. %d\n", omp_get_thread_num());
#pragma omp single nowait
        printf("Finishing work1. \n");
#pragma omp single nowait
        printf("Beginning work2. \n");
        printf("work on 2 parallelly. %d\n", omp_get_thread_num());
#pragma omp single
        printf("Finishing work2. \n");
    }
}


#endif //OPENMP_OPENMP_H
