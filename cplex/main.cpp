#include <vector>
#include <ilcplex/ilocplex.h>
ILOSTLBEGIN
//Expects an input of n, m, d values, constraint values
//Example: ./program.out 2 3 1 0 1 1 1 0 0.5 0.5 0.33333333337 0.33333333337 0.33333333337
int main (int argc, char **argv) {
  IloEnv env;
  int err = EXIT_SUCCESS;
  int n = atoi(argv[1]), m = atoi(argv[2]),
      arrayLength = n*m;

	std::streambuf *psbuf, *backup;
  std::ofstream filestr;
  filestr.open ("cplex.log");

  backup = std::clog.rdbuf();     // back up cout's streambuf

  psbuf = filestr.rdbuf();        // get file's streambuf
  std::clog.rdbuf(psbuf);         // assign streambuf to cout

  try {
        int i = 3, k = 0;

    IloModel model(env);

    clog << "bounds" << endl;
    //Define non-negative variables
    IloNumVarArray vars(env);
    for (int i = 0; i < arrayLength; i++) {
      vars.add(IloNumVar(env, 0.0, ILOFLOAT));
      clog << "x_" << i << " >= 0" << endl;
    }

    clog << "min" << endl;
    //Define the linear function to be minimized
    IloObjective obj = IloMinimize(env);
    while (i < arrayLength + 3) {
      clog << atof(argv[i]) <<  " x_" << k << " + ";
      obj.setLinearCoef(vars[k++], atof(argv[i++]));
    }
    clog << endl;
    model.add(obj);

    clog << "st" << endl;
    //Define problem constraints equal value
    IloRangeArray c(env);
    while (i < argc) {
      double constraint = atof(argv[i++]);
      clog << "c[" << i-10 << "] = " << constraint << endl;
      c.add(IloRange(env, constraint, constraint));
    }
    
    //Define coefficients at the form x_11 + x_12 + x_13
    k = 0;
    i = 0;
    while (i < n) {
      int j = 0;
      clog << "c[" << i << "]";
      while (j < m) {
        clog << " x_" << k << " +";
        c[i].setLinearCoef(vars[k++], 1.0);
        j++; 
      }
      clog << endl;
      i++;
    }

    //Define coefficients at the form x_11 + x_21 + x_31
    k = 0;
    while (i < n + m) {
      int j = 0;
      clog << "c[" << i << "]";
      while (j < n) {
        clog << " x_" << (j*m)+k << " +";
        c[i].setLinearCoef(vars[(j * m) + k], 1.0);
        j++; 
      }
      clog << endl;
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
    clog << "Solution status = " << cplex.getStatus() << endl;
    //clog << "Solution value = " << cplex.getObjValue() << endl;
    cplex.getValues(vals, vars);
    env.out() << "Values = " << vals << endl;
  }
  catch (IloException& e) {
    cerr << "Concert exception caught: " << e << endl;
    err = EXIT_FAILURE;
  }
  catch (...) {
    cerr << "Unknown exception caught" << endl;
    err = EXIT_FAILURE;
  }
  env.end();
  
  std::clog.rdbuf(backup); 
  filestr.close();
  return err;
}
