using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore.Metadata;

namespace FirstAPI.Models
{
    public partial class StudentsDbContext : DbContext
    {
        public StudentsDbContext()
        {
        }

        public StudentsDbContext(DbContextOptions<StudentsDbContext> options) : base(options)
        {
        }

        public DbSet<Students> Students { get; set; }

    }
}
