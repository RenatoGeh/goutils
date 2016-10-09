#include <stdio.h>
#include <stdlib.h>

#include <gsl/gsl_cdf.h>

int main(int argc, char *args[]) {
  int df;
  double chi;

  if (argc != 3) {
    printf("Usage: ./chisq chi df\n"
           "  chi - chi value for cumulative probability function\n"
           "  df  - degree of freedom\n");
    return 1;
  }

  chi = atof(args[1]);
  df = atoi(args[2]);

  printf("Pr(X^2 <= %f) = %f ~ X^2(%d)\n", chi, gsl_cdf_chisq_P(chi, df), df);

  return 0;
}
