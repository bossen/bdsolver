#include <vector>
#include <ilcplex/ilocplex.h>
ILOSTLBEGIN
//Expects an input of n, m, d values, constraint values
//Example: ./program.out 2 3 1 0 1 1 1 0 0.5 0.5 0.33333333337 0.33333333337 0.33333333337
int main (int argc, char **argv) {
  IloEnv env;
  int err = 0;
  int n = atoi(argv[1]), m = atoi(argv[2]),
      arrayLength = n*m;
  try {
        int i = 3, k = 0;

    IloModel model(env);

    //env.out() << "bounds" << endl;
    //Define non-negative variables
    IloNumVarArray vars(env);
    for (int i = 0; i < arrayLength; i++) {
      vars.add(IloNumVar(env, 0.0, ILOFLOAT));
      //env.out() << "x_" << i << " >= 0" << endl;
    }

    //env.out() << "min" << endl;
    //Define the linear function to be minimized
    IloObjective obj = IloMinimize(env);
    while (i < arrayLength + 3) {
      //env.out() << atof(argv[i]) <<  " x_" << k << " + ";
      obj.setLinearCoef(vars[k++], atof(argv[i++]));
    }
    //env.out() << endl;
    model.add(obj);

    //env.out() << "st" << endl;
    //Define problem constraints equal value
    IloRangeArray c(env);
    while (i < argc) {
      double constraint = atof(argv[i++]);
      //env.out() << "c[" << i-10 << "] = " << constraint << endl;
      c.add(IloRange(env, constraint, constraint));
    }
    
    //Define coefficients at the form x_11 + x_12 + x_13
    k = 0;
    i = 0;
    while (i < n) {
      int j = 0;
      //env.out() << "c[" << i << "]";
      while (j < m) {
        //env.out() << " x_" << k << " +";
        c[i].setLinearCoef(vars[k++], 1.0);
        j++; 
      }
      //env.out() << endl;
      i++;
    }

    //Define coefficients at the form x_11 + x_21 + x_31
    k = 0;
    while (i < n + m) {
      int j = 0;
      //env.out() << "c[" << i << "]";
      while (j < n) {
        //env.out() << " x_" << (j*m)+k << " +";
        c[i].setLinearCoef(vars[(j * m) + k], 1.0);
        j++; 
      }
      //env.out() << endl;
      k++;
      i++;
    }

    model.add(c);

    IloCplex cplex(model);
    cplex.setOut(env.getNullStream());
	
    if ( !cplex.solve() ) {
      env.error() << "Failed to optimize LP." << endl;
      throw(-1);
    }
    IloNumArray vals(env);
    //env.out() << "Solution status = " << cplex.getStatus() << endl;
    //env.out() << "Solution value = " << cplex.getObjValue() << endl;
    cplex.getValues(vals, vars);
    env.out() << "Values = " << vals << endl;
  }
  catch (IloException& e) {
    cerr << "Concert exception caught: " << e << endl;
    err = 1;
  }
  catch (...) {
    cerr << "Unknown exception caught" << endl;
    err = 1;
  }
  env.end();
  
  if (err) {
    return EXIT_FAILURE;
  }
  return EXIT_SUCCESS;
}
