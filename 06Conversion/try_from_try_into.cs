using System;
using System.Collections.Generic;
public class Program
{
    public class Result<T>
    {
        public T Value { get; private set; }
        public bool IsSuccess { get; private set; }
        public string ErrorMessage { get; private set; }
        public bool IsFailure => !IsSuccess;

        protected Result(T value, bool isSuccess, string errorMessage)
        {
            Value = value;
            IsSuccess = isSuccess;
            ErrorMessage = errorMessage;
        }

        public static Result<T> Ok(T value) => new Result<T>(value, true, null);
        public static Result<T> Fail(string message) => new Result<T>(default, false, message);

        public override bool Equals(object obj)
        {
            //这里的方式很好用 obj is Result<T> other 判断与赋值一起完成。
            return obj is Result<T> other&&
                EqualityComparer<T>.Default.Equals(Value, other.Value) &&
                IsSuccess == other.IsSuccess &&
                ErrorMessage == other.ErrorMessage;
        }

        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hash = 17;
                hash = hash * 23 + EqualityComparer<T>.Default.GetHashCode(Value);
                hash = hash * 23 + IsSuccess.GetHashCode();
                hash = hash * 23 + (ErrorMessage?.GetHashCode() ?? 0);
                return hash;
            }
        }

        public static bool operator ==(Result<T> left, Result<T> right)
        {
            return Equals(left, right);
        }

        public static bool operator !=(Result<T> left, Result<T> right)
        {
            return !Equals(left, right);
        }
    }
    public struct EvenNumber{
        public int value;
        public static  Result<EvenNumber> try_from( int value )
        {
            if( value % 2 ==  0 ){
                return Result<EvenNumber>.Ok( new EvenNumber{value = value} );
            }else{
                return Result<EvenNumber>.Fail($"{value} is not even.");
            }
        }
    }
    public static void Main(string[] args)
    {
        Console.WriteLine("This is C# Language");
        Console.WriteLine("EvenNumber.try_from(8) = Result<EvenNumber>.Ok( new EvenNumber{{value = 8}}): {0}" 
        ,EvenNumber.try_from(8) == Result<EvenNumber>.Ok( new EvenNumber{value = 8}));
        Console.WriteLine("EvenNumber.try_from(5) == Result<EvenNumber>.Fail(\"5 is not even.\"): {0}" 
        ,EvenNumber.try_from(5) == Result<EvenNumber>.Fail( "5 is not even."));
    }
}