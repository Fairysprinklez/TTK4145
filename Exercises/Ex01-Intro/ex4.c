#include <pthread.h>
#include <stdio.h>

int variable = 0;


void* threadFunc1(){
	for(int i=0; i<1000000; i++){
		 variable++;	
	}
	return NULL;
}

void* threadFunc2(){
        for(int i=0; i<1000000; i++){
                variable--;  
        }
        return NULL;
}

int main(){
	pthread_t thread1;
	pthread_t thread2;
	pthread_create(&thread1, NULL, threadFunc1, NULL);
	pthread_create(&thread2, NULL, threadFunc2, NULL);

	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);

	printf("%d\n", variable);
	
	return 0;
}
