//gcc -std=gnu99 -Wall -g -o ex4_fixed ex4_fixed.c -lpthread

#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>

sem_t mutex;

int variable = 0;


void* threadFunc1(){
	for(int i=0; i<1000000; i++){
		sem_wait(&mutex);
			variable++;
			//printf("%d\n", variable);
		sem_post(&mutex);	
	}
	return NULL;
}

void* threadFunc2(){
        for(int i=0; i<1000000; i++){
		sem_wait(&mutex);
                	variable--;
			//printf("%d\n", variable);
		sem_post(&mutex);  
        }
        return NULL;
}

int main(){

	sem_init(&mutex, 0, 1);	

	pthread_t thread1;
	pthread_t thread2;
	pthread_create(&thread1, NULL, threadFunc1, NULL);
	pthread_create(&thread2, NULL, threadFunc2, NULL);

	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);

	printf("%d\n", variable);
	
	return 0;
}
