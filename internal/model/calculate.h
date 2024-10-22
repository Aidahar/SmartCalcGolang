#ifndef SRC_CALCULATE_H_
#define SRC_CALCULATE_H_

//#include "math.h"
#include "parse_string.h"
#define MAXN 256

double calculate(char *notation, double x);
double take_number(char **notation);
int is_operator(char c);
int is_trig(char c);
void plus(double_node **numbers);
void minus(double_node **numbers);
void mul(double_node **numbers);
void divis(double_node **numbers);
void power(double_node **numbers);
void float_mod(double_node **numbers);
void doubl_sin(double_node **numbers);
void doubl_asin(double_node **numbers);
void doubl_cos(double_node **numbers);
void doubl_acos(double_node **numbers);
void doubl_tan(double_node **numbers);
void doubl_atan(double_node **numbers);
void doubl_ln(double_node **numbers);
void doubl_log(double_node **numbers);
void doubl_sqrt(double_node **numbers);

#endif  // SRC_CALCULATE_H
