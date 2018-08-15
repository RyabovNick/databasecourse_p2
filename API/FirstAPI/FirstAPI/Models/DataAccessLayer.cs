using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace FirstAPI.Models
{
    public class DataAccessLayer
    {
        StudentsDbContext db = new StudentsDbContext();

        public List<Students> GetStudents()
        {
            List<Students> lstStud = new List<Students>();
            lstStud = (from stud in db.Students select stud).ToList();

            return lstStud;
        }
    }
}
