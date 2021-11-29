class Kata
    {
        static public string ROT135(string text)
        {
            char[] vs = text.ToCharArray();

            for (int i = 0; i <= vs.Length - 1; i++)
            {
                if (vs[i] >= 65 && vs[i] <= 90)
                {
                    if (vs[i] >= 78)
                    {
                        vs[i] -= (char)13;
                    }
                    else
                        vs[i] += (char)13;
                }
                else if ((int)vs[i] >= 97 & (int)vs[i] <= 122)
                {
                    if (vs[i] >= 110)
                    {
                        vs[i] -= (char)13;
                    }
                    else
                        vs[i] += (char)13;
                }
            }

            for (int i = 0; i <= vs.Length - 1; i++)
            {
                if (vs[i] >= 48 && vs[i] <= 57)
                {
                    if (vs[i] >= 53)
                        vs[i] -= (char)5;
                    else
                        vs[i] += (char)5;
                }
            }

            System.Text.StringBuilder builder = new System.Text.StringBuilder();

            foreach (char value in vs)
            {
                builder.Append(value);
            }
            string result = builder.ToString();
            return result;
        }
    }